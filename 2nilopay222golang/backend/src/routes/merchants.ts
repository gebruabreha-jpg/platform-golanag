import { Router } from 'express';
import { db } from '../config/database';
import { validate } from '../utils/validator';
import { createMerchantSchema } from '../utils/validationSchemas';

const router = Router();

/**
 * GET /api/merchants
 * Get list of verified merchants (public)
 */
router.get('/', async (req, res) => {
  const { page = 1, limit = 20, businessType } = req.query;

  try {
    const skip = (Number(page) - 1) * Number(limit);
    const where: any = { isActive: true, isVerified: true };

    if (businessType) {
      where.businessType = businessType;
    }

    const [merchants, total] = await Promise.all([
      db.merchant.findMany({
        where,
        select: {
          id: true,
          name: true,
          businessType: true,
          categories: {
            select: { category: true },
          },
        },
        skip,
        take: Number(limit),
      }),
      db.merchant.count({ where }),
    ]);

    res.json({ merchants, total });
  } catch (error: any) {
    res.status(500).json({ error: 'Failed to fetch merchants' });
  }
});

/**
 * GET /api/merchants/:id
 * Get single merchant details
 */
router.get('/:id', async (req, res) => {
  try {
    const merchant = await db.merchant.findUnique({
      where: { id: req.params.id },
      select: {
        id: true,
        name: true,
        businessType: true,
        email: true,
        phone: true,
        address: true,
        categories: {
          select: { category: true },
        },
      },
    });

    if (!merchant || !merchant.isActive) {
      return res.status(404).json({ error: 'Merchant not found' });
    }

    res.json({ merchant });
  } catch (error: any) {
    res.status(500).json({ error: 'Failed to fetch merchant' });
  }
});

export default router;
