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
    'demo_user_00000000000000001',
    'demo@example.com',
    '$2a$12$LJ3m4ys3Lk0TSwHCpN1Stu/bJqHGey7p3BPLmOaVCtSbxDMbq7Ly',
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
    'demo_staff_0000000000000001',
    'staff-demo@example.com',
    '$2a$12$LJ3m4ys3Lk0TSwHCpN1Stu/bJqHGey7p3BPLmOaVCtSbxDMbq7Ly',
    'Demo Staff',
    'staff',
    TRUE, TRUE, TRUE,
    'email', 'active', 5000,
    TRUE, 'v1.0',
    NOW(), NOW()
) ON CONFLICT (id) DO NOTHING;
