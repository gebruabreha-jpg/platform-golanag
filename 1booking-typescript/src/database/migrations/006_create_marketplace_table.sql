CREATE TYPE item_condition AS ENUM ('NEW', 'LIKE_NEW', 'GOOD', 'FAIR');
CREATE TYPE payment_status AS ENUM ('PENDING', 'COMPLETED', 'CANCELLED', 'DISPUTED');

CREATE TABLE IF NOT EXISTS marketplace_items (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  seller_id UUID NOT NULL REFERENCES users(id),
  title VARCHAR(150) NOT NULL,
  description TEXT,
  category VARCHAR(80) NOT NULL,
  subcategory VARCHAR(80),
  price NUMERIC(12,2) NOT NULL,
  currency VARCHAR(3) NOT NULL DEFAULT 'USD',
  condition item_condition NOT NULL,
  location VARCHAR(120),
  country VARCHAR(80),
  shipping_available BOOLEAN NOT NULL DEFAULT FALSE,
  shipping_cost NUMERIC(10,2) NOT NULL DEFAULT 0,
  image_urls TEXT[],
  is_active BOOLEAN NOT NULL DEFAULT TRUE,
  is_sold BOOLEAN NOT NULL DEFAULT FALSE,
  view_count INTEGER NOT NULL DEFAULT 0,
  interest_count INTEGER NOT NULL DEFAULT 0,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_marketplace_seller ON marketplace_items(seller_id);
CREATE INDEX idx_marketplace_active ON marketplace_items(is_active, is_sold);
CREATE INDEX idx_marketplace_category ON marketplace_items(category);

CREATE TABLE IF NOT EXISTS transactions (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  buyer_id UUID NOT NULL REFERENCES users(id),
  seller_id UUID NOT NULL REFERENCES users(id),
  item_id UUID REFERENCES marketplace_items(id),
  type VARCHAR(30) NOT NULL,
  amount NUMERIC(12,2) NOT NULL,
  currency VARCHAR(3) NOT NULL DEFAULT 'USD',
  status payment_status NOT NULL DEFAULT 'PENDING',
  payment_method VARCHAR(50),
  escrow_id UUID,
  released_at TIMESTAMPTZ,
  completed_at TIMESTAMPTZ,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_transactions_buyer ON transactions(buyer_id);
CREATE INDEX idx_transactions_seller ON transactions(seller_id);
CREATE INDEX idx_transactions_status ON transactions(status);