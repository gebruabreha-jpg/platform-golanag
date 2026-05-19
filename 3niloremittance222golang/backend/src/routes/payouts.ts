import { Router } from 'express';
import { db } from '../config/database';
import { authenticateToken, AuthenticatedRequest } from '../middleware/auth';
import { wiseService } from '../services/wiseService';

const router = Router();

/**
 * GET /api/payouts
 * Get list of payouts (admin only or merchant own payouts)
 */
router.get('/', authenticateToken, async (req: AuthenticatedRequest, res) => {
  const { page = 1, limit = 20 } = req.query;
  const skip = (Number(page) - 1) * Number(limit);

  try {
    const where = {};

    // Non-admins can only see their own payouts
    if (req.user!.role !== 'PLATFORM_ADMIN') {
      // Join through transactions to filter by sender or merchant
      where.transaction = {
        OR: [
          { senderId: req.user!.id },
          { merchant: { userId: req.user!.id } },
        ],
      };
    }

    const [payouts, total] = await Promise.all([
      db.payout.findMany({
        where,
        include: {
          transaction: {
            include: {
              merchant: {
                select: { id: true, name: true },
              },
            },
          },
        },
        orderBy: { createdAt: 'desc' },
        skip,
        take: Number(limit),
      }),
      db.payout.count({ where }),
    ]);

    res.json({
      payouts,
      pagination: {
        page: Number(page),
        limit: Number(limit),
        total,
        pages: Math.ceil(total / Number(limit)),
      },
    });
  } catch (error: any) {
    res.status(500).json({ error: 'Failed to fetch payouts' });
  }
});

/**
 * GET /api/payouts/:id/status
 * Check payout status
 */
router.get('/:id/status', authenticateToken, async (req: AuthenticatedRequest, res) => {
  try {
    const payout = await db.payout.findUnique({
      where: { id: req.params.id },
      include: { transaction: true },
    });

    if (!payout) {
      return res.status(404).json({ error: 'Payout not found' });
    }

    // Check access
    if (
      payout.transaction.senderId !== req.user!.id &&
      req.user!.role !== 'PLATFORM_ADMIN'
    ) {
      return res.status(403).json({ error: 'Unauthorized' });
    }

    // Optionally fetch live status from Wise
    let wiseStatus;
    try {
      wiseStatus = await wiseService.getTransferStatus(payout.wiseTransferId);
    } catch (error) {
      // Wise API might be down, return stored status
      wiseStatus = payout.status;
    }

    res.json({
      payoutId: payout.id,
      status: wiseStatus || payout.status,
      createdAt: payout.createdAt,
      wiseTransferId: payout.wiseTransferId,
    });
  } catch (error: any) {
    res.status(500).json({ error: 'Failed to fetch payout status' });
  }
});

export default router;
