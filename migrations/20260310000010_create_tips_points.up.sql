CREATE TABLE IF NOT EXISTS tips (
  id VARCHAR(26) PRIMARY KEY,
  sender_id VARCHAR(26) NOT NULL REFERENCES users(id),
  recipient_id VARCHAR(26) NOT NULL REFERENCES users(id),
  staff_profile_id VARCHAR(26) NOT NULL REFERENCES staff_profiles(id),
  amount INTEGER NOT NULL,
  message TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS idx_tips_sender_id ON tips(sender_id);
CREATE INDEX IF NOT EXISTS idx_tips_recipient_id ON tips(recipient_id);

CREATE TABLE IF NOT EXISTS point_transactions (
  id VARCHAR(26) PRIMARY KEY,
  user_id VARCHAR(26) NOT NULL REFERENCES users(id),
  type VARCHAR(50) NOT NULL,
  amount INTEGER NOT NULL,
  balance_after INTEGER NOT NULL,
  reference_id VARCHAR(26),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS idx_point_transactions_user_id ON point_transactions(user_id);
