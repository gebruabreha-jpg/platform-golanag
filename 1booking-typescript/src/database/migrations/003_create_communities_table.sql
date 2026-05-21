CREATE TYPE community_category AS ENUM ('SHIPPING', 'HOUSING', 'MARKETPLACE', 'JOBS', 'SCHOLARSHIPS', 'BUSINESS');
CREATE TYPE post_type AS ENUM ('OFFER', 'REQUEST', 'INFO', 'DISCUSSION');

CREATE TABLE IF NOT EXISTS communities (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(120) NOT NULL,
  description TEXT,
  category community_category NOT NULL,
  location VARCHAR(120),
  country VARCHAR(80),
  is_private BOOLEAN NOT NULL DEFAULT FALSE,
  member_count INTEGER NOT NULL DEFAULT 0,
  moderator_id UUID NOT NULL REFERENCES users(id),
  rules TEXT,
  image_url TEXT,
  tags TEXT[],
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  CONSTRAINT chk_name CHECK (char_length(name) >= 2)
);

CREATE INDEX idx_communities_category ON communities(category);
CREATE INDEX idx_communities_moderator ON communities(moderator_id);
CREATE INDEX idx_communities_country ON communities(country);

CREATE TABLE IF NOT EXISTS community_members (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  community_id UUID NOT NULL REFERENCES communities(id) ON DELETE CASCADE,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  role VARCHAR(20) NOT NULL DEFAULT 'MEMBER',
  joined_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  UNIQUE(community_id, user_id)
);