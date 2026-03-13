CREATE TABLE IF NOT EXISTS reviews (
  id VARCHAR(26) PRIMARY KEY,
  booking_id VARCHAR(26) NOT NULL UNIQUE REFERENCES bookings(id),
  reviewer_id VARCHAR(26) NOT NULL REFERENCES users(id),
  reviewee_id VARCHAR(26) NOT NULL REFERENCES users(id),
  rating SMALLINT NOT NULL CHECK (rating >= 1 AND rating <= 5),
  comment TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS idx_reviews_reviewee_id ON reviews(reviewee_id);
