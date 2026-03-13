-- Rename unique constraints to match GORM's naming convention (uni_{table}_{column}).
-- Required because SQL migrations created constraints with PostgreSQL auto-naming,
-- but GORM AutoMigrate expects its own convention.

ALTER TABLE users RENAME CONSTRAINT users_email_key TO uni_users_email;
ALTER TABLE users RENAME CONSTRAINT users_google_id_key TO uni_users_google_id;
ALTER TABLE users RENAME CONSTRAINT users_apple_id_key TO uni_users_apple_id;
ALTER TABLE refresh_tokens RENAME CONSTRAINT refresh_tokens_token_hash_key TO uni_refresh_tokens_token_hash;
ALTER TABLE password_reset_tokens RENAME CONSTRAINT password_reset_tokens_token_hash_key TO uni_password_reset_tokens_token_hash;
ALTER TABLE staff_profiles RENAME CONSTRAINT staff_profiles_user_id_key TO uni_staff_profiles_user_id;
ALTER TABLE staff_profiles RENAME CONSTRAINT staff_profiles_staff_number_key TO uni_staff_profiles_staff_number;
ALTER TABLE reviews RENAME CONSTRAINT reviews_booking_id_key TO uni_reviews_booking_id;
