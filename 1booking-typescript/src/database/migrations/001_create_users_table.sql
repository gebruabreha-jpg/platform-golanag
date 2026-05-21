CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  email VARCHAR(255) NOT NULL UNIQUE,
  phone VARCHAR(32) UNIQUE,
  password_hash TEXT NOT NULL,
  first_name VARCHAR(80) NOT NULL DEFAULT '',
  last_name  VARCHAR(80) NOT NULL DEFAULT '',
  date_of_birth DATE,
  bio TEXT,
  avatar_url    TEXT,
  location      VARCHAR(120),
  country       VARCHAR(80),
  city          VARCHAR(80),
  role          VARCHAR(20) NOT NULL DEFAULT 'DIASPORA'
               CHECK (role IN ('DIASPORA','LOCAL','MERCHANT','ADMIN')),
  is_verified   BOOLEAN NOT NULL DEFAULT FALSE,
  verification_level INT NOT NULL DEFAULT 0,
  trust_score   DOUBLE PRECISION NOT NULL DEFAULT 0,
  total_transactions INTEGER NOT NULL DEFAULT 0,
  created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  deleted_at    TIMESTAMPTZ,
  CONSTRAINT email_valid CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$')
);

CREATE INDEX idx_users_email   ON users(email);
CREATE INDEX idx_users_role    ON users(role);
CREATE INDEX idx_users_deleted_at ON users(deleted_at) WHERE deleted_at IS NULL;