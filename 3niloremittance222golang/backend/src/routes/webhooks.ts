import { Router } from 'express';
import Stripe from 'stripe';
import { db } from '../config/database';
import { wiseService } from '../services/wiseService';
import { logger } from '../utils/logger';

const router = Router();

// Initialize Stripe webhook constructor
const stripeWebhook = new Stripe(process.env.STRIPE_SECRET_KEY!, {
  apiVersion: '2023-10-16',
});

/**
 * POST /api/webhooks/stripe
 * Handle Stripe webhook events
 */
router.post('/stripe', async (req, res) => {
  const sig = req.headers['stripe-signature'] as string;

  let event: Stripe.Event;

  try {
    event = stripeWebhook.webhooks.constructEvent(
      req.body,
      sig,
      process.env.STRIPE_WEBHOOK_SECRET!
    );
  } catch (err: any) {
    logger.error('Webhook signature verification failed', { error: err.message });
    return res.status(400).send(`Webhook Error: ${err.message}`);
  }

  try {
    switch (event.type) {
      case 'payment_intent.succeeded':
        await handlePaymentSucceeded(event.data.object as Stripe.PaymentIntent);
        break;

      case 'payment_intent.payment_failed':
        await handlePaymentFailed(event.data.object as Stripe.PaymentIntent);
        break;

      case 'charge.refunded':
        await handleRefund(event.data.object as Stripe.Charge);
        break;

      default:
        logger.info(`Unhandled webhook event type: ${event.type}`);
    }

    res.json({ received: true });
  } catch (error: any) {
    logger.error('Webhook processing error', {
      type: event.type,
      error: error.message,
    });
    res.status(500).json({ error: 'Webhook handler failed' });
  }
});

/**
 * Handle successful payment
 */
async function handlePaymentSucceeded(paymentIntent: Stripe.PaymentIntent) {
  const { id: stripePaymentId, metadata } = paymentIntent;

  if (!metadata?.referenceId) {
    logger.error('Payment intent missing referenceId', { stripePaymentId });
    return;
  }

  const transaction = await db.transaction.findFirst({
    where: { stripePaymentId },
  });

  if (!transaction) {
    logger.error('Transaction not found for payment', { stripePaymentId });
    return;
  }

  // Update transaction status
  await db.transaction.update({
    where: { id: transaction.id },
    data: {
      status: 'PROCESSING',
      stripeChargeId: paymentIntent.charges.data[0]?.id,
    },
  });

  // Initiate payout to merchant via Wise
  await wiseService.createPayout(transaction);

  logger.info('Payment succeeded, payout initiated', {
    transactionId: transaction.id,
    referenceId: transaction.referenceId,
  });
}

/**
 * Handle failed payment
 */
async function handlePaymentFailed(paymentIntent: Stripe.PaymentIntent) {
  const { id: stripePaymentId } = paymentIntent;

  const transaction = await db.transaction.findFirst({
    where: { stripePaymentId },
  });

  if (!transaction) {
    logger.error('Transaction not found for failed payment', { stripePaymentId });
    return;
  }

  await db.transaction.update({
    where: { id: transaction.id },
    data: {
      status: 'FAILED',
      failureReason: paymentIntent.last_payment_error?.message || 'Payment failed',
    },
  });

  logger.warn('Payment failed', {
    transactionId: transaction.id,
    referenceId: transaction.referenceId,
    reason: paymentIntent.last_payment_error?.message,
  });
}

/**
 * Handle refund
 */
async function handleRefund(charge: Stripe.Charge) {
  const { payment_intent: paymentIntentId } = charge;

  const transaction = await db.transaction.findFirst({
    where: { stripePaymentId: paymentIntentId as string },
  });

  if (!transaction) {
    logger.error('Transaction not found for refund', { paymentIntentId });
    return;
  }

  await db.transaction.update({
    where: { id: transaction.id },
    data: { status: 'REFUNDED' },
  });

  logger.info('Transaction refunded', {
    transactionId: transaction.id,
    referenceId: transaction.referenceId,
  });
}

export default router;
