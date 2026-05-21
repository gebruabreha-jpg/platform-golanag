import Joi from 'joi';

export const registerSchema = Joi.object({
  email: Joi.string().email().required(),
  password: Joi.string().min(8).required(),
  firstName: Joi.string().min(2).max(50).required(),
  lastName: Joi.string().min(2).max(50).required(),
  role: Joi.string().valid('DIASPORA_USER', 'MERCHANT_ADMIN').optional(),
  phone: Joi.string().optional(),
});

export const loginSchema = Joi.object({
  email: Joi.string().email().required(),
  password: Joi.string().required(),
});

export const createTransactionSchema = Joi.object({
  merchantId: Joi.string().required(),
  amount: Joi.number().positive().precision(2).required(),
  serviceType: Joi.string().valid('TUITION', 'HEALTHCARE', 'GROCERY', 'UTILITY').optional(),
  description: Joi.string().max(500).optional(),
  metadata: Joi.object().optional(),
});

export const createMerchantSchema = Joi.object({
  name: Joi.string().min(2).max(100).required(),
  businessType: Joi.string().valid('SCHOOL', 'HOSPITAL', 'SUPERMARKET', 'PHARMACY', 'UTILITY_COMPANY').required(),
  registrationNumber: Joi.string().optional(),
  taxId: Joi.string().optional(),
  email: Joi.string().email().optional(),
  phone: Joi.string().optional(),
  bankName: Joi.string().required(),
  bankAccountName: Joi.string().required(),
  bankAccountNumber: Joi.string().required(),
  branchName: Joi.string().optional(),
  address: Joi.object().optional(),
});

/**
 * Generic validation helper
 */
export function validate(schema: Joi.Schema, data: any) {
  return schema.validate(data, { abortEarly: false });
}
