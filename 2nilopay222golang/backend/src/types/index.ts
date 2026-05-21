export interface User {
  id: string;
  email: string;
  phone?: string;
  firstName: string;
  lastName: string;
  role: 'DIASPORA_USER' | 'MERCHANT_ADMIN' | 'PLATFORM_ADMIN';
  isActive: boolean;
  isVerified: boolean;
  createdAt: string;
  updatedAt?: string;
}

export interface Merchant {
  id: string;
  name: string;
  businessType: 'SCHOOL' | 'HOSPITAL' | 'SUPERMARKET' | 'PHARMACY' | 'UTILITY_COMPANY';
  registrationNumber?: string;
  taxId?: string;
  email?: string;
  phone?: string;
  bankName: string;
  bankAccountName: string;
  bankAccountNumber: string;
  isActive: boolean;
  isVerified: boolean;
}

export interface Transaction {
  id: string;
  referenceId: string;
  senderId: string;
  merchantId: string;
  amount: string; // Stored as string for precision
  localAmount: string;
  feeAmount: string;
  totalCharged: string;
  fxRate: string;
  status: 'PENDING' | 'PROCESSING' | 'COMPLETED' | 'FAILED' | 'REFUNDED' | 'CANCELLED';
  stripePaymentId?: string;
  wiseTransferId?: string;
  failureReason?: string;
  completedAt?: string;
  createdAt: string;
  updatedAt: string;
  merchant?: {
    id: string;
    name: string;
    businessType: string;
  };
}

export interface CreateTransactionRequest {
  merchantId: string;
  amount: number;
  serviceType?: string;
  description?: string;
  metadata?: Record<string, any>;
}

export interface Payout {
  id: string;
  transactionId: string;
  wiseTransferId: string;
  amount: string;
  localAmount: string;
  feeAmount: string;
  fxRate: string;
  status: 'PENDING' | 'PROCESSING' | 'COMPLETED' | 'FAILED';
  createdAt: string;
}
