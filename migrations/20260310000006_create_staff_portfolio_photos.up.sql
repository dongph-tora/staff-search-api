CREATE TABLE IF NOT EXISTS staff_portfolio_photos (
  id VARCHAR(26) PRIMARY KEY,
  staff_profile_id VARCHAR(26) NOT NULL REFERENCES staff_profiles(id) ON DELETE CASCADE,
  photo_url TEXT NOT NULL,
  display_order INTEGER NOT NULL DEFAULT 0,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS idx_portfolio_staff_profile_id ON staff_portfolio_photos(staff_profile_id);
