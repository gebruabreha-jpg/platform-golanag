import Stripe from 'stripe';
import { logger } from '../utils/logger';

export const stripe = new Stripe(process.env.STRIPE_SECRET_KEY!, {
  apiVersion: '2023-10-16',
});

/**
 * Create a payment intent for a transaction
 */
export async function createPaymentIntent(
  amount: number,
  currency: string = 'usd',
  metadata: Record<string, string> = {}
) {
  try {
    const paymentIntent = await stripe.paymentIntents.create({
      amount: Math.round(amount * 100), // Convert to cents
      currency,
      automatic_payment_methods: { enabled: true },
      metadata,
      description: `Nilo Commerce Payment - ${metadata.serviceType || 'Bill Payment'}`,
    });

    logger.info('Payment intent created', { paymentIntentId: paymentIntent.id });
    return paymentIntent;
  } catch (error: any) {
    logger.error('Failed to create payment intent', { error: error.message });
    throw error;
  }
}

/**
 * Retrieve payment intent by ID
 */
export async function getPaymentIntent(paymentIntentId: string) {
  try {
    return await stripe.paymentIntents.retrieve(paymentIntentId);
  } catch (error: any) {
    logger.error('Failed to retrieve payment intent', { error: error.message });
    throw error;
  }
}

/**
 * Create a refund for a charge
 */
export async function createRefund(chargeId: string, amount?: number) {
  try {
    const refund = await stripe.refunds.create({
      charge: chargeId,
      amount: amount ? Math.round(amount * 100) : undefined,
    });

    logger.info('Refund created', { refundId: refund.id, chargeId });
    return refund;
  } catch (error: any) {
    logger.error('Failed to create refund', { error: error.message });
    throw error;
  }
}

/**
 * Verify webhook signature
 */
export function constructEvent(payload: string | Buffer, signature: string) {
  return stripe.webhooks.constructEvent(payload, signature, process.env.STRIPE_WEBHOOK_SECRET!);
}
