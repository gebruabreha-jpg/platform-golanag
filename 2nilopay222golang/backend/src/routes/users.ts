import { Router } from 'express';
import { db } from '../config/database';
import { authenticateToken, AuthenticatedRequest } from '../middleware/auth';

const router = Router();

/**
 * GET /api/users/profile
 * Get user profile
 */
router.get('/profile', authenticateToken, async (req: AuthenticatedRequest, res) => {
  try {
    const user = await db.user.findUnique({
      where: { id: req.user!.id },
      select: {
        id: true,
        email: true,
        phone: true,
        firstName: true,
        lastName: true,
        role: true,
        isVerified: true,
        verifiedAt: true,
        createdAt: true,
      },
    });

    res.json({ user });
  } catch (error: any) {
    res.status(500).json({ error: 'Failed to fetch profile' });
  }
});

/**
 * PUT /api/users/profile
 * Update user profile
 */
router.put('/profile', authenticateToken, async (req: AuthenticatedRequest, res) => {
  const { firstName, lastName, phone } = req.body;

  try {
    const user = await db.user.update({
      where: { id: req.user!.id },
      data: {
        firstName,
        lastName,
        phone,
      },
      select: {
        id: true,
        email: true,
        firstName: true,
        lastName: true,
        phone: true,
      },
    });

    res.json({ user });
  } catch (error: any) {
    res.status(500).json({ error: 'Failed to update profile' });
  }
});

export default router;
