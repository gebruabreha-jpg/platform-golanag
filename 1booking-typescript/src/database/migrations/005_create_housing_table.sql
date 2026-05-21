CREATE TYPE property_type AS ENUM ('APARTMENT', 'HOUSE', 'ROOM', 'STUDIO');
CREATE TYPE room_type AS ENUM ('PRIVATE', 'SHARED', 'MASTER');
CREATE TYPE lease_term AS ENUM ('MONTHLY', 'SHORT_TERM', 'LONG_TERM');

CREATE TABLE IF NOT EXISTS housing_listings (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  landlord_id UUID NOT NULL REFERENCES users(id),
  title VARCHAR(150) NOT NULL,
  description TEXT,
  property_type property_type NOT NULL,
  room_type room_type,
  bedrooms INTEGER NOT NULL DEFAULT 1,
  bathrooms INTEGER NOT NULL DEFAULT 1,
  monthly_rent NUMERIC(12,2) NOT NULL,
  deposit NUMERIC(12,2),
  address VARCHAR(200) NOT NULL,
  city VARCHAR(80) NOT NULL,
  country VARCHAR(80) NOT NULL,
  latitude DOUBLE PRECISION NOT NULL,
  longitude DOUBLE PRECISION NOT NULL,
  available_from DATE,
  lease_term lease_term NOT NULL,
  furnished BOOLEAN NOT NULL DEFAULT FALSE,
  includes_utilities BOOLEAN NOT NULL DEFAULT FALSE,
  image_urls TEXT[],
  is_active BOOLEAN NOT NULL DEFAULT TRUE,
  application_count INTEGER NOT NULL DEFAULT 0,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_housing_landlord ON housing_listings(landlord_id);
CREATE INDEX idx_housing_active ON housing_listings(is_active);
CREATE INDEX idx_housing_location ON housing_listings(country, city);

CREATE TABLE IF NOT EXISTS housing_applications (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  listing_id UUID NOT NULL REFERENCES housing_listings(id) ON DELETE CASCADE,
  user_id UUID NOT NULL REFERENCES users(id),
  message TEXT,
  move_in_date DATE,
  proposed_rent NUMERIC(12,2),
  status VARCHAR(20) NOT NULL DEFAULT 'PENDING',
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  UNIQUE(listing_id, user_id)
);