CREATE TABLE IF NOT EXISTS staff_profiles (
  id VARCHAR(26) PRIMARY KEY,
  user_id VARCHAR(26) NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
  staff_number VARCHAR(6) NOT NULL UNIQUE,
  job_title VARCHAR(100) NOT NULL,
  job_category VARCHAR(50) NOT NULL,
  location VARCHAR(255),
  bio TEXT,
  intro_video_url TEXT,
  is_available BOOLEAN NOT NULL DEFAULT FALSE,
  accept_bookings BOOLEAN NOT NULL DEFAULT TRUE,
  rating DECIMAL(3,2) NOT NULL DEFAULT 0.00,
  review_count INTEGER NOT NULL DEFAULT 0,
  followers_count INTEGER NOT NULL DEFAULT 0,
  total_tips_received INTEGER NOT NULL DEFAULT 0,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS idx_staff_profiles_user_id ON staff_profiles(user_id);
CREATE INDEX IF NOT EXISTS idx_staff_profiles_job_category ON staff_profiles(job_category);
CREATE INDEX IF NOT EXISTS idx_staff_profiles_is_available ON staff_profiles(is_available);
