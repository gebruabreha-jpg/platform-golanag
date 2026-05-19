export interface User {
  id: string;
  email: string;
  first_name: string;
  last_name: string;
  phone?: string;
  avatar_url?: string;
  location?: string;
  country?: string;
  city?: string;
  role: 'DIASPORA' | 'LOCAL' | 'MERCHANT' | 'ADMIN';
  is_verified: boolean;
  verification_level: number;
  trust_score: number;
  total_transactions: number;
  created_at: string;
  updated_at?: string;
}

export interface Community {
  id: string;
  name: string;
  description: string;
  category: 'SHIPPING' | 'HOUSING' | 'MARKETPLACE' | 'JOBS' | 'SCHOLARSHIPS' | 'BUSINESS';
  location: string;
  country: string;
  is_private: boolean;
  member_count: number;
  moderator_id: string;
  rules?: string;
  image_url?: string;
  tags?: string[];
  created_at: string;
  updated_at: string;
}

export interface Post {
  id: string;
  community_id: string;
  user_id: string;
  type: 'OFFER' | 'REQUEST' | 'INFO' | 'DISCUSSION';
  title: string;
  content: string;
  media_urls?: string[];
  is_pinned: boolean;
  is_closed: boolean;
  reply_count: number;
  view_count: number;
  created_at: string;
  updated_at: string;
}

export interface MarketplaceItem {
  id: string;
  seller_id: string;
  title: string;
  description: string;
  category: string;
  subcategory?: string;
  price: number;
  currency: 'USD' | 'ETB' | 'EUR' | 'GBP';
  condition: 'NEW' | 'LIKE_NEW' | 'GOOD' | 'FAIR';
  location: string;
  country: string;
  shipping_available: boolean;
  shipping_cost: number;
  image_urls?: string[];
  is_active: boolean;
  is_sold: boolean;
  view_count: number;
  interest_count: number;
  created_at: string;
  updated_at: string;
}

export interface HousingListing {
  id: string;
  landlord_id: string;
  title: string;
  description: string;
  property_type: 'APARTMENT' | 'HOUSE' | 'ROOM';
  room_type: 'PRIVATE' | 'SHARED' | 'MASTER';
  bedrooms: number;
  bathrooms: number;
  monthly_rent: number;
  currency: string;
  deposit: number;
  address: string;
  city: string;
  country: string;
  latitude?: number;
  longitude?: number;
  available_from: string;
  lease_term: 'MONTHLY' | 'SHORT_TERM' | 'LONG_TERM';
  furnished: boolean;
  includes_utilities: boolean;
  image_urls?: string[];
  is_active: boolean;
  application_count: number;
  created_at: string;
  updated_at: string;
}

export interface Scholarship {
  id: string;
  title: string;
  description: string;
  provider: string;
  provider_type: 'GOVERNMENT' | 'UNIVERSITY' | 'FOUNDATION';
  country: string;
  city?: string;
  level: 'UNDERGRADUATE' | 'GRADUATE' | 'PHD' | 'POSTDOC';
  field: string;
  amount: number;
  currency: string;
  covers: string[];
  deadline?: string;
  eligibility: string;
  requirements: string;
  application_url: string;
  is_active: boolean;
  is_featured: boolean;
}

export interface Job {
  id: string;
  employer_id: string;
  title: string;
  description: string;
  job_type: 'FULL_TIME' | 'PART_TIME' | 'CONTRACT' | 'FREELANCE';
  remote: boolean;
  location: string;
  country: string;
  salary_min: number;
  salary_max: number;
  currency: string;
  industry: string;
  skills: string[];
  benefits?: string[];
  application_url: string;
  expires_at?: string;
  is_active: boolean;
  view_count: number;
  application_count: number;
}

export interface Transaction {
  id: string;
  buyer_id: string;
  seller_id: string;
  item_id?: string;
  type: 'MARKETPLACE_PURCHASE' | 'SERVICE_PAYMENT' | 'DONATION';
  amount: number;
  currency: string;
  status: 'PENDING' | 'COMPLETED' | 'CANCELLED' | 'DISPUTED';
  payment_method: string;
  escrow_id?: string;
  released_at?: string;
  completed_at?: string;
  created_at: string;
}

export interface CurrencyRate {
  id: string;
  from_currency: string;
  to_currency: string;
  rate: number;
  source: string;
  updated_at: string;
}

export interface VerificationResult {
  is_authentic: boolean;
  authenticity_score: number;
  risk_flags: string[];
  recommendations: string[];
  confidence: string;
}

export interface HousingMatch {
  listing_id: string;
  score: number;
  match_factors: Record<string, number>;
  reasons: string[];
}

export interface AIRecommendation {
  item_id: string;
  score: number;
  reason: string;
  category: string;
}
