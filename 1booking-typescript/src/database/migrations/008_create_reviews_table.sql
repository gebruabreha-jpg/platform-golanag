CREATE TYPE subject_type AS ENUM ('USER', 'LISTING', 'COMMUNITY');

CREATE TABLE IF NOT EXISTS trust_reviews (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  reviewer_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  subject_id UUID NOT NULL,
  subject_type subject_type NOT NULL,
  rating INTEGER NOT NULL CHECK (rating >= 1 AND rating <= 5),
  comment TEXT,
  is_verified BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  CONSTRAINT chk_review_subject CHECK (subject_id IS NOT NULL)
);

CREATE INDEX idx_reviews_reviewer ON trust_reviews(reviewer_id);
CREATE INDEX idx_reviews_subject ON trust_reviews(subject_id, subject_type);
CREATE INDEX idx_reviews_verified ON trust_reviews(is_verified);

CREATE TABLE IF NOT EXISTS reviews (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  booking_id UUID,
  reviewer_id UUID NOT NULL REFERENCES users(id),
  reviewed_id UUID NOT NULL REFERENCES users(id),
  rating INTEGER NOT NULL CHECK (rating >= 1 AND rating <= 5),
  comment TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);