import { Router } from 'express';
import { db } from '../config/database';
import { stripe } from '../services/stripeService';
import { wiseService } from '../services/wiseService';
import { authenticateToken, AuthenticatedRequest } from '../middleware/auth';
import { logger } from '../utils/logger';
import { v4 as uuidv4 } from 'uuid';

const router = Router();

/**
 * GET /api/transactions
 * Get list of user's transactions
 */
router.get('/', authenticateToken, async (req: AuthenticatedRequest, res) => {
  try {
    const { page = 1, limit = 20, status } = req.query;
    const skip = (Number(page) - 1) * Number(limit);

    const where: any = {
      OR: [
        { senderId: req.user!.id },
        // merchants could see their own transactions via different endpoint
      ],
    };

    if (status) {
      where.status = status;
    }

    const [transactions, total] = await Promise.all([
      db.transaction.findMany({
        where,
        include: {
          merchant: {
            select: {
              id: true,
              name: true,
              businessType: true,
            },
          },
        },
        orderBy: { createdAt: 'desc' },
        skip,
        take: Number(limit),
      }),
      db.transaction.count({ where }),
    ]);

    res.json({
      transactions,
      pagination: {
        page: Number(page),
        limit: Number(limit),
        total,
        pages: Math.ceil(total / Number(limit)),
      },
    });
  } catch (error: any) {
    logger.error('Get transactions error', { error: error.message });
    res.status(500).json({ error: 'Failed to fetch transactions' });
  }
});

/**
 * GET /api/transactions/:id
 * Get single transaction details
 */
router.get('/:id', authenticateToken, async (req: AuthenticatedRequest, res) => {
  try {
    const { id } = req.params;

    const transaction = await db.transaction.findUnique({
      where: { id },
      include: {
        merchant: {
          select: {
            id: true,
            name: true,
            businessType: true,
            bankName: true,
          },
        },
        payout: true,
      },
    });

    if (!transaction) {
      return res.status(404).json({ error: 'Transaction not found' });
    }

    // Ensure user owns this transaction or is platform admin
    if (transaction.senderId !== req.user!.id && req.user!.role !== 'PLATFORM_ADMIN') {
      return res.status(403).json({ error: 'Unauthorized' });
    }

    res.json({ transaction });
  } catch (error: any) {
    logger.error('Get transaction error', { error: error.message });
    res.status(500).json({ error: 'Failed to fetch transaction' });
  }
});

/**
 * POST /api/transactions
 * Create a new bill payment transaction
 */
router.post('/', authenticateToken, async (req: AuthenticatedRequest, res) => {
  try {
    const { merchantId, amount, serviceType, description, metadata } = req.body;

    // Validate input
    if (!merchantId || !amount || amount <= 0) {
      return res.status(400).json({
        error: 'Merchant ID and valid amount are required',
      });
    }

    // Verify merchant exists and is active
    const merchant = await db.merchant.findUnique({
      where: { id: merchantId },
    });

    if (!merchant || !merchant.isActive) {
      return res.status(404).json({ error: 'Merchant not found or inactive' });
    }

    // Generate unique reference ID for diaspora user to track
    const referenceId = `NLC-${Date.now()}-${uuidv4().slice(0, 8).toUpperCase()}`;

    // Calculate fee and total
    const feePercentage = parseFloat(process.env.DEFAULT_FEE_PERCENTAGE || '2.5');
    const feeAmount = (amount * feePercentage) / 100;
    const totalCharged = amount + feeAmount;

    // Get current FX rate from Wise (or use cached rate)
    const fxRate = await wiseService.getExchangeRate('USD', 'ETB');
    const localAmount = amount * fxRate;

    // Create Stripe payment intent
    const paymentIntent = await stripe.paymentIntents.create({
      amount: Math.round(totalCharged * 100), // Convert to cents
      currency: 'usd',
      description: `Payment to ${merchant.name} - ${serviceType || 'Bill Payment'}`,
      metadata: {
        referenceId,
        merchantId,
        senderId: req.user!.id,
        serviceType: serviceType || '',
      },
    });

    // Create transaction record
    const transaction = await db.transaction.create({
      data: {
        referenceId,
        senderId: req.user!.id,
        merchantId,
        amount: amount.toString(),
        fxRate: fxRate.toString(),
        localAmount: localAmount.toString(),
        feeAmount: feeAmount.toString(),
        totalCharged: totalCharged.toString(),
        stripePaymentId: paymentIntent.id,
        status: 'PENDING',
      },
      include: {
        merchant: true,
      },
    });

    logger.info('Transaction created', {
      transactionId: transaction.id,
      referenceId,
      amount,
      merchantId,
    });

    res.status(201).json({
      message: 'Transaction created',
      transaction,
      paymentIntent: {
        clientSecret: paymentIntent.client_secret,
      },
    });
  } catch (error: any) {
    logger.error('Create transaction error', { error: error.message });
    res.status(500).json({ error: 'Failed to create transaction' });
  }
});

/**
 * POST /api/transactions/:id/confirm
 * Confirm payment completion (used by webhook from Stripe)
 */
router.post('/:id/confirm', async (req, res) => {
  const { id } = req.params;

  try {
    const transaction = await db.transaction.findUnique({
      where: { id },
    });

    if (!transaction) {
      return res.status(404).json({ error: 'Transaction not found' });
    }

    // Get updated payment status from Stripe
    const paymentIntent = await stripe.paymentIntents.retrieve(
      transaction.stripePaymentId!
    );

    if (paymentIntent.status === 'succeeded') {
      // Update transaction status
      await db.transaction.update({
        where: { id },
        data: {
          status: 'PROCESSING',
          stripeChargeId: paymentIntent.charges.data[0]?.id,
        },
      });

      // Trigger payout to merchant via Wise
      await wiseService.createPayout(transaction);

      logger.info('Payment confirmed, payout initiated', { transactionId: id });
    }

    res.json({ message: 'Payment status updated' });
  } catch (error: any) {
    logger.error('Confirm transaction error', { error: error.message });
    res.status(500).json({ error: 'Failed to confirm transaction' });
  }
});

export default router;
