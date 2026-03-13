CREATE TABLE IF NOT EXISTS services (
  id VARCHAR(26) PRIMARY KEY,
  staff_profile_id VARCHAR(26) NOT NULL REFERENCES staff_profiles(id) ON DELETE CASCADE,
  name VARCHAR(255) NOT NULL,
  description TEXT,
  price DECIMAL(10,2) NOT NULL,
  duration_minutes INTEGER NOT NULL,
  is_active BOOLEAN NOT NULL DEFAULT TRUE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS bookings (
  id VARCHAR(26) PRIMARY KEY,
  user_id VARCHAR(26) NOT NULL REFERENCES users(id),
  staff_profile_id VARCHAR(26) NOT NULL REFERENCES staff_profiles(id),
  service_id VARCHAR(26) REFERENCES services(id),
  status VARCHAR(20) NOT NULL DEFAULT 'pending',
  note TEXT,
  scheduled_at TIMESTAMPTZ,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS idx_bookings_user_id ON bookings(user_id);
CREATE INDEX IF NOT EXISTS idx_bookings_staff_profile_id ON bookings(staff_profile_id);
CREATE INDEX IF NOT EXISTS idx_bookings_status ON bookings(status);
