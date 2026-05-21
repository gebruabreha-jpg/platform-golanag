import axios from 'axios';
import { logger } from '../utils/logger';

const WISE_API_BASE = 'https://api.wise.com';
const WISE_API_VERSION = 'v3';

const wiseClient = axios.create({
  baseURL: `${WISE_API_BASE}/${WISE_API_VERSION}`,
  headers: {
    Authorization: `Bearer ${process.env.WISE_API_TOKEN}`,
    Content-Type: 'application/json',
  },
});

/**
 * Get exchange rate between two currencies
 */
export async function getExchangeRate(sourceCurrency: string, targetCurrency: string): Promise<number> {
  try {
    const response = await wiseClient.get('/rates', {
      params: {
        source: sourceCurrency,
        target: targetCurrency,
      },
    });

    const rate = response.data[0]?.rate;
    if (!rate) {
      throw new Error('Exchange rate not available');
    }

    logger.info('FX rate retrieved', { source: sourceCurrency, target: targetCurrency, rate });
    return rate;
  } catch (error: any) {
    logger.error('Failed to get exchange rate', { error: error.message });
    // Fallback to a reasonable rate for development
    if (process.env.NODE_ENV === 'development') {
      return 55.5; // Approximate USD/ETB rate (update as needed)
    }
    throw error;
  }
}

/**
 * Create a payout to an Ethiopian merchant
 */
export async function createPayout(transaction: any): Promise<any> {
  try {
    const profileId = process.env.WISE_API_PROFILE_ID;
    if (!profileId) {
      throw new Error('Wise profile ID not configured');
    }

    // Get merchant bank details
    const merchant = await db.merchant.findUnique({
      where: { id: transaction.merchantId },
    });

    if (!merchant) {
      throw new Error('Merchant not found');
    }

    // Convert amount to ETB
    const localAmount = parseFloat(transaction.localAmount);
    const fxRate = parseFloat(transaction.fxRate);

    // Prepare recipient
    const recipient = await ensureRecipient(merchant);

    // Create transfer
    const transferResponse = await wiseClient.post(`/profiles/${profileId}/transfers`, {
      targetAccount: recipient.id,
      sourceCurrency: 'USD',
      targetCurrency: 'ETB',
      amount: localAmount,
      reference: `Payment for ${merchant.name} - ${transaction.referenceId}`,
      customerReferenceId: transaction.referenceId,
    });

    const transfer = transferResponse.data;

    // Create payout record
    await db.payout.create({
      data: {
        transactionId: transaction.id,
        wiseTransferId: transfer.id,
        wiseProfileId: profileId,
        amount: transaction.amount,
        localAmount: localAmount.toString(),
        feeAmount: '0', // Wise fee included in spread
        fxRate: fxRate.toString(),
        status: 'PROCESSING',
      },
    });

    // Update transaction status
    await db.transaction.update({
      where: { id: transaction.id },
      data: {
        status: 'PROCESSING',
        wiseTransferId: transfer.id,
      },
    });

    logger.info('Payout created', {
      transactionId: transaction.id,
      wiseTransferId: transfer.id,
      amount: localAmount,
    });

    return transfer;
  } catch (error: any) {
    logger.error('Failed to create payout', { error: error.message });
    throw error;
  }
}

/**
 * Ensure recipient exists in Wise, create if not
 */
async function ensureRecipient(merchant: any): Promise<{ id: string }> {
  try {
    // Search for existing recipient by bank account
    const response = await wiseClient.get('/accounts', {
      params: {
        currency: 'ETB',
      },
    });

    // This is simplified - implement actual recipient lookup/creation
    // based on Wise API documentation for your specific use case
    const account = response.data.accounts[0];
    return { id: account?.id || 'placeholder-account-id' };
  } catch (error: any) {
    logger.error('Failed to ensure recipient', { error: error.message });
    throw error;
  }
}

/**
 * Check transfer status
 */
export async function getTransferStatus(transferId: string): Promise<string> {
  try {
    const response = await wiseClient.get(`/transfers/${transferId}`);
    return response.data.status;
  } catch (error: any) {
    logger.error('Failed to get transfer status', { error: error.message });
    throw error;
  }
}

/**
 * Cancel a pending transfer
 */
export async function cancelTransfer(transferId: string): Promise<void> {
  try {
    await wiseClient.post(`/transfers/${transferId}/cancel`);
    logger.info('Transfer cancelled', { transferId });
  } catch (error: any) {
    logger.error('Failed to cancel transfer', { error: error.message });
    throw error;
  }
}

export const wiseService = {
  getExchangeRate,
  createPayout,
  getTransferStatus,
  cancelTransfer,
};
