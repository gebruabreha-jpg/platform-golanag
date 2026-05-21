CREATE TYPE provider_type AS ENUM ('GOVERNMENT', 'UNIVERSITY', 'FOUNDATION');
CREATE TYPE job_type AS ENUM ('FULL_TIME', 'PART_TIME', 'CONTRACT', 'FREELANCE');

CREATE TABLE IF NOT EXISTS scholarships (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  title VARCHAR(200) NOT NULL,
  description TEXT NOT NULL,
  provider VARCHAR(120) NOT NULL,
  provider_type provider_type NOT NULL,
  country VARCHAR(80) NOT NULL,
  city VARCHAR(80),
  level VARCHAR(50) NOT NULL,
  field VARCHAR(80) NOT NULL,
  amount NUMERIC(12,2),
  currency VARCHAR(3),
  is_active BOOLEAN NOT NULL DEFAULT TRUE,
  is_featured BOOLEAN NOT NULL DEFAULT FALSE,
  deadline DATE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_scholarships_active ON scholarships(is_active);
CREATE INDEX idx_scholarships_featured ON scholarships(is_featured);
CREATE INDEX idx_scholarships_country ON scholarships(country);

CREATE TABLE IF NOT EXISTS jobs (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  employer_id UUID NOT NULL REFERENCES users(id),
  title VARCHAR(150) NOT NULL,
  description TEXT NOT NULL,
  job_type job_type NOT NULL,
  remote BOOLEAN NOT NULL DEFAULT FALSE,
  location VARCHAR(120),
  country VARCHAR(80),
  salary_min NUMERIC(12,2),
  salary_max NUMERIC(12,2),
  currency VARCHAR(3),
  industry VARCHAR(80),
  skills TEXT[],
  benefits TEXT[],
  application_url TEXT,
  expires_at TIMESTAMPTZ,
  is_active BOOLEAN NOT NULL DEFAULT TRUE,
  view_count INTEGER NOT NULL DEFAULT 0,
  application_count INTEGER NOT NULL DEFAULT 0,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_jobs_employer ON jobs(employer_id);
CREATE INDEX idx_jobs_active ON jobs(is_active);
CREATE INDEX idx_jobs_country ON jobs(country);