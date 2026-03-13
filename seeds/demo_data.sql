-- Demo seed data — run only on development/staging environments.
-- Safe to run multiple times (ON CONFLICT DO NOTHING).
--
-- Passwords: both accounts use "demo123" (bcrypt cost 12).
-- Run: psql $DATABASE_URL -f seeds/demo_data.sql

INSERT INTO users (
    id, email, password_hash, name, role,
    is_staff, is_staff_registered, is_verified,
    auth_provider, status, points,
    privacy_policy_accepted, privacy_policy_version,
    created_at, updated_at
) VALUES (
    'demo_user_0000000000000001',
    'demo@example.com',
    '$2a$12$Jy9j1BJRuwktBv9YCMmgRO/Tj7kOJ4i1JJvBjBMEZ.LiDxuGwpWdC',
    'Demo User',
    'user',
    FALSE, FALSE, TRUE,
    'email', 'active', 1000,
    TRUE, 'v1.0',
    NOW(), NOW()
) ON CONFLICT (id) DO NOTHING;

INSERT INTO users (
    id, email, password_hash, name, role,
    is_staff, is_staff_registered, is_verified,
    auth_provider, status, points,
    privacy_policy_accepted, privacy_policy_version,
    created_at, updated_at
) VALUES (
    'demo_staff_000000000000001',
    'staff-demo@example.com',
    '$2a$12$Jy9j1BJRuwktBv9YCMmgRO/Tj7kOJ4i1JJvBjBMEZ.LiDxuGwpWdC',
    'Demo Staff',
    'staff',
    TRUE, TRUE, TRUE,
    'email', 'active', 5000,
    TRUE, 'v1.0',
    NOW(), NOW()
) ON CONFLICT (id) DO NOTHING;

INSERT INTO staff_profiles (
    id, user_id, staff_number,
    job_title, job_category,
    location, bio,
    is_available, accept_bookings,
    rating, review_count, followers_count, total_tips_received,
    created_at, updated_at
) VALUES (
    'demo_staff_prof_000000001',
    'demo_staff_000000000000001',
    '000001',
    'ネイリスト',
    'nail_art',
    '東京都渋谷区',
    'プロのネイリストです。お気軽にご予約ください。',
    TRUE, TRUE,
    4.8, 12, 34, 15000,
    NOW(), NOW()
) ON CONFLICT (id) DO NOTHING;
