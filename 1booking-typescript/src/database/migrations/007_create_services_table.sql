CREATE TYPE service_category AS ENUM ('LEGAL', 'MEDICAL', 'CONSULTING', 'TECHNICAL', 'CREATIVE', 'TRADES', 'TRANSPORT', 'OTHER');
CREATE TYPE lawyer_specialization AS ENUM ('IMMIGRATION', 'BUSINESS', 'FAMILY', 'PROPERTY', 'CRIMINAL');

CREATE TABLE IF NOT EXISTS service_professionals (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  specialization VARCHAR(50) NOT NULL,
  name VARCHAR(120) NOT NULL,
  firm VARCHAR(120),
  years_experience INTEGER NOT NULL DEFAULT 0,
  location VARCHAR(120) NOT NULL,
  country VARCHAR(80) NOT NULL,
  consultation_fee NUMERIC(10,2) NOT NULL,
  currency VARCHAR(3) NOT NULL DEFAULT 'USD',
  is_verified BOOLEAN NOT NULL DEFAULT FALSE,
  rating DOUBLE PRECISION NOT NULL DEFAULT 0,
  review_count INTEGER NOT NULL DEFAULT 0,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_services_specialization ON service_professionals(specialization);
CREATE INDEX idx_services_location ON service_professionals(country, location);

CREATE TABLE IF NOT EXISTS lawyer_bookings (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  lawyer_id UUID NOT NULL REFERENCES service_professionals(id) ON DELETE CASCADE,
  user_id UUID NOT NULL REFERENCES users(id),
  scheduled_at TIMESTAMPTZ NOT NULL,
  duration INTEGER NOT NULL,
  type VARCHAR(20) NOT NULL CHECK (type IN ('IN_PERSON', 'REMOTE')),
  status VARCHAR(20) NOT NULL DEFAULT 'PENDING',
  notes TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_bookings_lawyer ON lawyer_bookings(lawyer_id);
CREATE INDEX idx_bookings_user ON lawyer_bookings(user_id);
CREATE INDEX idx_bookings_scheduled ON lawyer_bookings(scheduled_at);