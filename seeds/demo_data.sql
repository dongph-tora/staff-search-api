-- Demo seed data — run only on development/staging environments.
-- Safe to run multiple times (ON CONFLICT DO NOTHING).
--
-- Passwords: all accounts use "demo123" (bcrypt cost 12).
-- Run: psql $DATABASE_URL -f seeds/demo_data.sql

-- ============================================================
-- USERS — Customers
-- ============================================================

INSERT INTO users (
    id, email, password_hash, name, phone_number, avatar_url, role,
    is_staff, is_staff_registered, is_verified,
    auth_provider, status, points,
    privacy_policy_accepted, privacy_policy_version,
    created_at, updated_at
) VALUES
(
    'demo_user_0000000000000001',
    'demo@example.com',
    '$2a$12$Jy9j1BJRuwktBv9YCMmgRO/Tj7kOJ4i1JJvBjBMEZ.LiDxuGwpWdC',
    'Demo User', NULL, 'https://i.pravatar.cc/300?img=3',
    'user', FALSE, FALSE, TRUE,
    'email', 'active', 1000,
    TRUE, 'v1.0',
    NOW() - INTERVAL '30 days', NOW()
),
(
    'seed_user_0000000000000001',
    'yuki.tanaka@example.com',
    '$2a$12$Jy9j1BJRuwktBv9YCMmgRO/Tj7kOJ4i1JJvBjBMEZ.LiDxuGwpWdC',
    'Yuki Tanaka', '090-1234-5678', 'https://i.pravatar.cc/300?img=1',
    'user', FALSE, FALSE, TRUE,
    'email', 'active', 500,
    TRUE, 'v1.0',
    NOW() - INTERVAL '25 days', NOW()
),
(
    'seed_user_0000000000000002',
    'hana.yamamoto@example.com',
    '$2a$12$Jy9j1BJRuwktBv9YCMmgRO/Tj7kOJ4i1JJvBjBMEZ.LiDxuGwpWdC',
    'Hana Yamamoto', '080-2345-6789', 'https://i.pravatar.cc/300?img=5',
    'user', FALSE, FALSE, TRUE,
    'email', 'active', 200,
    TRUE, 'v1.0',
    NOW() - INTERVAL '20 days', NOW()
),
(
    'seed_user_0000000000000003',
    'kenji.sato@example.com',
    '$2a$12$Jy9j1BJRuwktBv9YCMmgRO/Tj7kOJ4i1JJvBjBMEZ.LiDxuGwpWdC',
    'Kenji Sato', '070-3456-7890', 'https://i.pravatar.cc/300?img=60',
    'user', FALSE, FALSE, TRUE,
    'email', 'active', 1500,
    TRUE, 'v1.0',
    NOW() - INTERVAL '15 days', NOW()
),
(
    'seed_user_0000000000000004',
    'mei.suzuki@example.com',
    '$2a$12$Jy9j1BJRuwktBv9YCMmgRO/Tj7kOJ4i1JJvBjBMEZ.LiDxuGwpWdC',
    'Mei Suzuki', '090-4567-8901', 'https://i.pravatar.cc/300?img=9',
    'user', FALSE, FALSE, TRUE,
    'email', 'active', 800,
    TRUE, 'v1.0',
    NOW() - INTERVAL '10 days', NOW()
)
ON CONFLICT (id) DO NOTHING;

-- ============================================================
-- USERS — Staff members
-- ============================================================

INSERT INTO users (
    id, email, password_hash, name, phone_number, avatar_url, role,
    is_staff, is_staff_registered, is_verified,
    auth_provider, status, points,
    privacy_policy_accepted, privacy_policy_version,
    created_at, updated_at
) VALUES
(
    'demo_staff_000000000000001',
    'staff-demo@example.com',
    '$2a$12$Jy9j1BJRuwktBv9YCMmgRO/Tj7kOJ4i1JJvBjBMEZ.LiDxuGwpWdC',
    'Demo Staff', NULL, 'https://i.pravatar.cc/300?img=16',
    'staff', TRUE, TRUE, TRUE,
    'email', 'active', 5000,
    TRUE, 'v1.0',
    NOW() - INTERVAL '60 days', NOW()
),
(
    'seed_staff_000000000000002',
    'sakura.ito@example.com',
    '$2a$12$Jy9j1BJRuwktBv9YCMmgRO/Tj7kOJ4i1JJvBjBMEZ.LiDxuGwpWdC',
    'Sakura Ito', '090-1111-2222', 'https://i.pravatar.cc/300?img=12',
    'staff', TRUE, TRUE, TRUE,
    'email', 'active', 8200,
    TRUE, 'v1.0',
    NOW() - INTERVAL '55 days', NOW()
),
(
    'seed_staff_000000000000003',
    'miku.watanabe@example.com',
    '$2a$12$Jy9j1BJRuwktBv9YCMmgRO/Tj7kOJ4i1JJvBjBMEZ.LiDxuGwpWdC',
    'Miku Watanabe', '080-2222-3333', 'https://i.pravatar.cc/300?img=20',
    'staff', TRUE, TRUE, TRUE,
    'email', 'active', 6500,
    TRUE, 'v1.0',
    NOW() - INTERVAL '50 days', NOW()
),
(
    'seed_staff_000000000000004',
    'yuna.kimura@example.com',
    '$2a$12$Jy9j1BJRuwktBv9YCMmgRO/Tj7kOJ4i1JJvBjBMEZ.LiDxuGwpWdC',
    'Yuna Kimura', '070-3333-4444', 'https://i.pravatar.cc/300?img=25',
    'staff', TRUE, TRUE, TRUE,
    'email', 'active', 9100,
    TRUE, 'v1.0',
    NOW() - INTERVAL '45 days', NOW()
),
(
    'seed_staff_000000000000005',
    'rina.hayashi@example.com',
    '$2a$12$Jy9j1BJRuwktBv9YCMmgRO/Tj7kOJ4i1JJvBjBMEZ.LiDxuGwpWdC',
    'Rina Hayashi', '090-4444-5555', 'https://i.pravatar.cc/300?img=29',
    'staff', TRUE, TRUE, TRUE,
    'email', 'active', 4300,
    TRUE, 'v1.0',
    NOW() - INTERVAL '40 days', NOW()
),
(
    'seed_staff_000000000000006',
    'haruto.yoshida@example.com',
    '$2a$12$Jy9j1BJRuwktBv9YCMmgRO/Tj7kOJ4i1JJvBjBMEZ.LiDxuGwpWdC',
    'Haruto Yoshida', '080-5555-6666', 'https://i.pravatar.cc/300?img=52',
    'staff', TRUE, TRUE, TRUE,
    'email', 'active', 3800,
    TRUE, 'v1.0',
    NOW() - INTERVAL '35 days', NOW()
),
(
    'seed_staff_000000000000007',
    'aoi.nakamura@example.com',
    '$2a$12$Jy9j1BJRuwktBv9YCMmgRO/Tj7kOJ4i1JJvBjBMEZ.LiDxuGwpWdC',
    'Aoi Nakamura', '070-6666-7777', 'https://i.pravatar.cc/300?img=32',
    'staff', TRUE, TRUE, TRUE,
    'email', 'active', 11500,
    TRUE, 'v1.0',
    NOW() - INTERVAL '30 days', NOW()
)
ON CONFLICT (id) DO NOTHING;

-- ============================================================
-- STAFF PROFILES
-- ============================================================

INSERT INTO staff_profiles (
    id, user_id, staff_number,
    job_title, job_category,
    location, latitude, longitude, bio,
    is_available, accept_bookings,
    rating, review_count, followers_count, total_tips_received,
    created_at, updated_at
) VALUES
(
    'demo_staff_prof_000000001',
    'demo_staff_000000000000001',
    '000001',
    'ネイリスト', 'nail_art',
    '東京都渋谷区', 35.6595, 139.7005,
    'プロのネイリスト歴8年。トレンドデザインからシンプルなものまで幅広く対応します。丁寧な施術でお客様に喜んでいただけるよう努めています。',
    TRUE, TRUE,
    4.8, 47, 128, 85000,
    NOW() - INTERVAL '60 days', NOW()
),
(
    'seed_prof_0000000000000001',
    'seed_staff_000000000000002',
    '000002',
    'ヘアスタイリスト', 'hair_stylist',
    '東京都新宿区', 35.6938, 139.7034,
    '美容師歴10年。カットからカラー、パーマまで何でもお任せください。お客様の骨格や髪質に合ったスタイルを提案します。毎月の講習で最新トレンドをキャッチ。',
    TRUE, TRUE,
    4.9, 63, 245, 142000,
    NOW() - INTERVAL '55 days', NOW()
),
(
    'seed_prof_0000000000000002',
    'seed_staff_000000000000003',
    '000003',
    'ネイリスト', 'nail_art',
    '大阪府大阪市北区', 34.7024, 135.4959,
    'ジェルネイル専門サロン出身。3Dアートや繊細なラインアートが得意です。SNSでも作品を公開中。お気に入りデザインの持ち込みもOK！',
    TRUE, TRUE,
    4.7, 38, 189, 67000,
    NOW() - INTERVAL '50 days', NOW()
),
(
    'seed_prof_0000000000000003',
    'seed_staff_000000000000004',
    '000004',
    'マッサージセラピスト', 'massage',
    '神奈川県横浜市中区', 35.4478, 139.6425,
    'タイ古式マッサージ資格保有。スウェーデン式・アロマトリートメントも対応。肩こり・腰痛でお悩みの方をはじめ、リラックスしたい方もぜひ。出張施術も可。',
    TRUE, TRUE,
    4.8, 55, 167, 98000,
    NOW() - INTERVAL '45 days', NOW()
),
(
    'seed_prof_0000000000000004',
    'seed_staff_000000000000005',
    '000005',
    'メイクアップアーティスト', 'makeup',
    '東京都港区', 35.6581, 139.7514,
    'ブライダル・成人式・撮影などのスペシャルオケージョンから日常メイクレッスンまで対応。元大手メイクサロン勤務。ナチュラル〜華やか幅広くお任せを。',
    FALSE, TRUE,
    4.6, 29, 312, 53000,
    NOW() - INTERVAL '40 days', NOW()
),
(
    'seed_prof_0000000000000005',
    'seed_staff_000000000000006',
    '000006',
    '理容師・バーバー', 'barber',
    '東京都渋谷区', 35.6580, 139.7016,
    'NY仕込みのバーバーテクニック。フェードカット・スキンフェード・ドレッドヘアが得意。シェービングも本格対応。男性のお客様大歓迎。',
    TRUE, TRUE,
    4.5, 41, 98, 44000,
    NOW() - INTERVAL '35 days', NOW()
),
(
    'seed_prof_0000000000000006',
    'seed_staff_000000000000007',
    '000007',
    'まつ毛スタイリスト', 'eyelash',
    '東京都目黒区', 35.6337, 139.7160,
    'まつ毛エクステ歴6年。ナチュラル〜ボリュームラッシュまで全対応。まつ毛パーマ・リフトアップも大好評。素材にこだわった安心安全な施術をお届けします。',
    TRUE, TRUE,
    4.9, 71, 421, 185000,
    NOW() - INTERVAL '30 days', NOW()
)
ON CONFLICT (id) DO NOTHING;

-- ============================================================
-- STAFF PORTFOLIO PHOTOS
-- ============================================================

INSERT INTO staff_portfolio_photos (id, staff_profile_id, photo_url, display_order, created_at) VALUES
-- Demo Staff (nail_art)
('seed_phot_0000000000000001', 'demo_staff_prof_000000001', 'https://picsum.photos/seed/nail1/400/500', 0, NOW() - INTERVAL '50 days'),
('seed_phot_0000000000000002', 'demo_staff_prof_000000001', 'https://picsum.photos/seed/nail2/400/500', 1, NOW() - INTERVAL '40 days'),
('seed_phot_0000000000000003', 'demo_staff_prof_000000001', 'https://picsum.photos/seed/nail3/400/500', 2, NOW() - INTERVAL '30 days'),
-- Sakura Ito (hair_stylist)
('seed_phot_0000000000000004', 'seed_prof_0000000000000001', 'https://picsum.photos/seed/hair1/400/500', 0, NOW() - INTERVAL '50 days'),
('seed_phot_0000000000000005', 'seed_prof_0000000000000001', 'https://picsum.photos/seed/hair2/400/500', 1, NOW() - INTERVAL '40 days'),
('seed_phot_0000000000000006', 'seed_prof_0000000000000001', 'https://picsum.photos/seed/hair3/400/500', 2, NOW() - INTERVAL '30 days'),
-- Miku Watanabe (nail_art)
('seed_phot_0000000000000007', 'seed_prof_0000000000000002', 'https://picsum.photos/seed/nail4/400/500', 0, NOW() - INTERVAL '45 days'),
('seed_phot_0000000000000008', 'seed_prof_0000000000000002', 'https://picsum.photos/seed/nail5/400/500', 1, NOW() - INTERVAL '35 days'),
('seed_phot_0000000000000009', 'seed_prof_0000000000000002', 'https://picsum.photos/seed/nail6/400/500', 2, NOW() - INTERVAL '25 days'),
-- Yuna Kimura (massage)
('seed_phot_0000000000000010', 'seed_prof_0000000000000003', 'https://picsum.photos/seed/spa1/400/500', 0, NOW() - INTERVAL '40 days'),
('seed_phot_0000000000000011', 'seed_prof_0000000000000003', 'https://picsum.photos/seed/spa2/400/500', 1, NOW() - INTERVAL '30 days'),
('seed_phot_0000000000000012', 'seed_prof_0000000000000003', 'https://picsum.photos/seed/spa3/400/500', 2, NOW() - INTERVAL '20 days'),
-- Rina Hayashi (makeup)
('seed_phot_0000000000000013', 'seed_prof_0000000000000004', 'https://picsum.photos/seed/makeup1/400/500', 0, NOW() - INTERVAL '35 days'),
('seed_phot_0000000000000014', 'seed_prof_0000000000000004', 'https://picsum.photos/seed/makeup2/400/500', 1, NOW() - INTERVAL '25 days'),
('seed_phot_0000000000000015', 'seed_prof_0000000000000004', 'https://picsum.photos/seed/makeup3/400/500', 2, NOW() - INTERVAL '15 days'),
-- Haruto Yoshida (barber)
('seed_phot_0000000000000016', 'seed_prof_0000000000000005', 'https://picsum.photos/seed/barber1/400/500', 0, NOW() - INTERVAL '30 days'),
('seed_phot_0000000000000017', 'seed_prof_0000000000000005', 'https://picsum.photos/seed/barber2/400/500', 1, NOW() - INTERVAL '20 days'),
('seed_phot_0000000000000018', 'seed_prof_0000000000000005', 'https://picsum.photos/seed/barber3/400/500', 2, NOW() - INTERVAL '10 days'),
-- Aoi Nakamura (eyelash)
('seed_phot_0000000000000019', 'seed_prof_0000000000000006', 'https://picsum.photos/seed/lash1/400/500', 0, NOW() - INTERVAL '28 days'),
('seed_phot_0000000000000020', 'seed_prof_0000000000000006', 'https://picsum.photos/seed/lash2/400/500', 1, NOW() - INTERVAL '18 days'),
('seed_phot_0000000000000021', 'seed_prof_0000000000000006', 'https://picsum.photos/seed/lash3/400/500', 2, NOW() - INTERVAL '8 days')
ON CONFLICT (id) DO NOTHING;

-- ============================================================
-- SERVICES
-- ============================================================

INSERT INTO services (id, staff_profile_id, name, description, price, duration_minutes, is_active, created_at, updated_at) VALUES
-- Demo Staff (nail_art)
('seed_srvc_0000000000000001', 'demo_staff_prof_000000001', 'ジェルネイル（シンプル）', 'シンプルカラーやグラデーションなどベーシックなジェルネイル', 5000.00, 60, TRUE, NOW() - INTERVAL '55 days', NOW()),
('seed_srvc_0000000000000002', 'demo_staff_prof_000000001', 'ジェルネイル（アート入り）', 'ストーン・ラメ・ネイルアートを含むデザインコース', 7500.00, 90, TRUE, NOW() - INTERVAL '55 days', NOW()),
('seed_srvc_0000000000000003', 'demo_staff_prof_000000001', 'ネイルオフ＋付け替え', '既存ネイルのオフから新しいデザインへの付け替え', 9000.00, 120, TRUE, NOW() - INTERVAL '55 days', NOW()),

-- Sakura Ito (hair_stylist)
('seed_srvc_0000000000000004', 'seed_prof_0000000000000001', 'カット', 'シャンプー・ブロー込みのスタンダードカット', 4500.00, 60, TRUE, NOW() - INTERVAL '50 days', NOW()),
('seed_srvc_0000000000000005', 'seed_prof_0000000000000001', 'カット＋カラー', 'カット＋全体カラーのセットメニュー', 12000.00, 150, TRUE, NOW() - INTERVAL '50 days', NOW()),
('seed_srvc_0000000000000006', 'seed_prof_0000000000000001', 'カット＋パーマ', 'カット込みのデジタルパーマまたは通常パーマ', 15000.00, 180, TRUE, NOW() - INTERVAL '50 days', NOW()),

-- Miku Watanabe (nail_art)
('seed_srvc_0000000000000007', 'seed_prof_0000000000000002', 'ジェルネイル（ワンカラー）', 'ワンカラーのシンプルジェルネイル', 4800.00, 60, TRUE, NOW() - INTERVAL '45 days', NOW()),
('seed_srvc_0000000000000008', 'seed_prof_0000000000000002', '3Dネイルアート', '立体的な3Dアートデザイン', 10000.00, 120, TRUE, NOW() - INTERVAL '45 days', NOW()),
('seed_srvc_0000000000000009', 'seed_prof_0000000000000002', 'ネイルオフのみ', 'ジェルネイルのオフのみ', 2500.00, 30, TRUE, NOW() - INTERVAL '45 days', NOW()),

-- Yuna Kimura (massage)
('seed_srvc_0000000000000010', 'seed_prof_0000000000000003', '全身マッサージ 60分', 'スウェーデン式全身リラクゼーションマッサージ', 8000.00, 60, TRUE, NOW() - INTERVAL '40 days', NOW()),
('seed_srvc_0000000000000011', 'seed_prof_0000000000000003', '全身マッサージ 90分', 'アロマオイルを使ったディープリラックスコース', 11000.00, 90, TRUE, NOW() - INTERVAL '40 days', NOW()),
('seed_srvc_0000000000000012', 'seed_prof_0000000000000003', 'ヘッドスパ＋肩甲骨はがし', '頭皮ケア＋肩周りの集中ケア', 6500.00, 60, TRUE, NOW() - INTERVAL '40 days', NOW()),

-- Rina Hayashi (makeup)
('seed_srvc_0000000000000013', 'seed_prof_0000000000000004', 'ブライダルメイク', '結婚式・前撮り対応のフルメイク（ヘアセット込）', 25000.00, 120, TRUE, NOW() - INTERVAL '35 days', NOW()),
('seed_srvc_0000000000000014', 'seed_prof_0000000000000004', 'パーティーメイク', '二次会・パーティー向けメイクアップ', 9000.00, 75, TRUE, NOW() - INTERVAL '35 days', NOW()),
('seed_srvc_0000000000000015', 'seed_prof_0000000000000004', 'メイクレッスン（1時間）', 'セルフメイクのスキルアップレッスン', 6000.00, 60, TRUE, NOW() - INTERVAL '35 days', NOW()),

-- Haruto Yoshida (barber)
('seed_srvc_0000000000000016', 'seed_prof_0000000000000005', 'カット', 'スタンダードバーバーカット', 3500.00, 45, TRUE, NOW() - INTERVAL '30 days', NOW()),
('seed_srvc_0000000000000017', 'seed_prof_0000000000000005', 'カット＋シェービング', 'カット＋本格レザーシェービング', 5500.00, 75, TRUE, NOW() - INTERVAL '30 days', NOW()),
('seed_srvc_0000000000000018', 'seed_prof_0000000000000005', 'フェードカット', 'スキンフェード・ハイフェード対応', 4500.00, 60, TRUE, NOW() - INTERVAL '30 days', NOW()),

-- Aoi Nakamura (eyelash)
('seed_srvc_0000000000000019', 'seed_prof_0000000000000006', 'まつ毛エクステ（80本）', 'ナチュラルに仕上げるシングルエクステ80本', 6500.00, 75, TRUE, NOW() - INTERVAL '25 days', NOW()),
('seed_srvc_0000000000000020', 'seed_prof_0000000000000006', 'まつ毛エクステ（120本）', 'ボリュームラッシュ仕上げ120本', 9500.00, 90, TRUE, NOW() - INTERVAL '25 days', NOW()),
('seed_srvc_0000000000000021', 'seed_prof_0000000000000006', 'まつ毛パーマ＋トリートメント', '自まつ毛を活かしたまつ毛パーマ', 7000.00, 60, TRUE, NOW() - INTERVAL '25 days', NOW())
ON CONFLICT (id) DO NOTHING;

-- ============================================================
-- BOOKINGS
-- ============================================================

INSERT INTO bookings (id, user_id, staff_profile_id, service_id, status, note, scheduled_at, created_at, updated_at) VALUES
-- Completed bookings (for review seeding)
('seed_book_0000000000000001', 'demo_user_0000000000000001', 'demo_staff_prof_000000001', 'seed_srvc_0000000000000002', 'completed',
 'アートデザインはお任せします', NOW() - INTERVAL '20 days', NOW() - INTERVAL '22 days', NOW() - INTERVAL '20 days'),
('seed_book_0000000000000002', 'seed_user_0000000000000001', 'seed_prof_0000000000000001', 'seed_srvc_0000000000000005', 'completed',
 'アッシュ系のカラーにしてください', NOW() - INTERVAL '18 days', NOW() - INTERVAL '20 days', NOW() - INTERVAL '18 days'),
('seed_book_0000000000000003', 'seed_user_0000000000000002', 'seed_prof_0000000000000003', 'seed_srvc_0000000000000011', 'completed',
 '肩と首周りが特に凝っています', NOW() - INTERVAL '14 days', NOW() - INTERVAL '15 days', NOW() - INTERVAL '14 days'),
('seed_book_0000000000000004', 'seed_user_0000000000000003', 'seed_prof_0000000000000006', 'seed_srvc_0000000000000020', 'completed',
 'ナチュラルに見えるボリュームでお願いします', NOW() - INTERVAL '10 days', NOW() - INTERVAL '11 days', NOW() - INTERVAL '10 days'),
('seed_book_0000000000000005', 'seed_user_0000000000000004', 'seed_prof_0000000000000005', 'seed_srvc_0000000000000017', 'completed',
 'スキンフェードでお願いします', NOW() - INTERVAL '8 days', NOW() - INTERVAL '9 days', NOW() - INTERVAL '8 days'),
('seed_book_0000000000000006', 'demo_user_0000000000000001', 'seed_prof_0000000000000001', 'seed_srvc_0000000000000004', 'completed',
 NULL, NOW() - INTERVAL '6 days', NOW() - INTERVAL '7 days', NOW() - INTERVAL '6 days'),

-- Confirmed bookings (upcoming)
('seed_book_0000000000000007', 'seed_user_0000000000000001', 'seed_prof_0000000000000006', 'seed_srvc_0000000000000019', 'confirmed',
 '初めてです。ナチュラルにお願いします', NOW() + INTERVAL '2 days', NOW() - INTERVAL '1 day', NOW()),
('seed_book_0000000000000008', 'seed_user_0000000000000002', 'demo_staff_prof_000000001', 'seed_srvc_0000000000000001', 'confirmed',
 'ピンク系でお願いします', NOW() + INTERVAL '3 days', NOW() - INTERVAL '1 day', NOW()),
('seed_book_0000000000000009', 'seed_user_0000000000000003', 'seed_prof_0000000000000001', 'seed_srvc_0000000000000006', 'confirmed',
 'ゆるふわパーマ希望', NOW() + INTERVAL '5 days', NOW() - INTERVAL '2 days', NOW()),

-- Pending bookings (awaiting confirmation)
('seed_book_0000000000000010', 'seed_user_0000000000000004', 'seed_prof_0000000000000004', 'seed_srvc_0000000000000014', 'pending',
 '友人の誕生日パーティー用です', NOW() + INTERVAL '7 days', NOW() - INTERVAL '3 hours', NOW()),
('seed_book_0000000000000011', 'demo_user_0000000000000001', 'seed_prof_0000000000000003', 'seed_srvc_0000000000000010', 'pending',
 NULL, NOW() + INTERVAL '4 days', NOW() - INTERVAL '1 hour', NOW()),

-- Cancelled bookings
('seed_book_0000000000000012', 'seed_user_0000000000000002', 'seed_prof_0000000000000002', 'seed_srvc_0000000000000008', 'cancelled',
 '急用が入ってしまいました。また予約します', NOW() - INTERVAL '3 days', NOW() - INTERVAL '5 days', NOW() - INTERVAL '4 days')
ON CONFLICT (id) DO NOTHING;

-- ============================================================
-- REVIEWS (for completed bookings only)
-- ============================================================

INSERT INTO reviews (id, booking_id, reviewer_id, reviewee_id, rating, comment, created_at) VALUES
(
    'seed_rvew_0000000000000001',
    'seed_book_0000000000000001',
    'demo_user_0000000000000001',
    'demo_staff_000000000000001',
    5,
    'デザインがとても可愛くて大満足です！丁寧な施術で痛みも全くなく、仕上がりも長持ちしています。また絶対来ます！',
    NOW() - INTERVAL '19 days'
),
(
    'seed_rvew_0000000000000002',
    'seed_book_0000000000000002',
    'seed_user_0000000000000001',
    'seed_staff_000000000000002',
    5,
    'アッシュカラーがイメージ通りに仕上がりました！骨格に合わせたカットの提案もしてくれて、とても頼りになるスタイリストさんです。',
    NOW() - INTERVAL '17 days'
),
(
    'seed_rvew_0000000000000003',
    'seed_book_0000000000000003',
    'seed_user_0000000000000002',
    'seed_staff_000000000000004',
    5,
    '肩こりが完全にほぐれました！アロマの香りも心地よく、90分があっという間でした。定期的に通いたいと思います。',
    NOW() - INTERVAL '13 days'
),
(
    'seed_rvew_0000000000000004',
    'seed_book_0000000000000004',
    'seed_user_0000000000000003',
    'seed_staff_000000000000007',
    5,
    '120本なのにとてもナチュラルに仕上げていただきました。持ちも良く、技術力の高さを感じます。接客も丁寧で居心地が良かったです。',
    NOW() - INTERVAL '9 days'
),
(
    'seed_rvew_0000000000000005',
    'seed_book_0000000000000005',
    'seed_user_0000000000000004',
    'seed_staff_000000000000006',
    4,
    'スキンフェードがきれいに決まりました！シェービングも初めてでしたが、とても気持ちよかったです。次回はカラーも試してみたいです。',
    NOW() - INTERVAL '7 days'
),
(
    'seed_rvew_0000000000000006',
    'seed_book_0000000000000006',
    'demo_user_0000000000000001',
    'seed_staff_000000000000002',
    5,
    'カットが上手でシャンプーも気持ちよかったです。仕上がりのブローも完璧でした。',
    NOW() - INTERVAL '5 days'
)
ON CONFLICT (id) DO NOTHING;

-- ============================================================
-- POSTS (feed content from staff)
-- ============================================================

INSERT INTO posts (id, author_id, content, media_url, media_type, likes_count, comments_count, is_active, created_at, updated_at) VALUES
-- Demo Staff posts
('seed_post_0000000000000001', 'demo_staff_000000000000001',
 '今日の作品 🌸 春らしいピンクとゴールドのグラデーションネイル。お客様にとても喜んでいただけました！ご予約はプロフィールから✨',
 'https://picsum.photos/seed/post_nail_a/600/600', 'image', 47, 8, TRUE, NOW() - INTERVAL '3 days', NOW()),
('seed_post_0000000000000002', 'demo_staff_000000000000001',
 '週末限定！フレンチネイル＋ストーンアート1本サービス中 💅 3月末までの予約限定です。お見逃しなく！',
 'https://picsum.photos/seed/post_nail_b/600/600', 'image', 62, 12, TRUE, NOW() - INTERVAL '1 day', NOW()),
('seed_post_0000000000000003', 'demo_staff_000000000000001',
 'ネイルケアのコツ💡 毎日就寝前にキューティクルオイルを塗ることでネイルの持ちが全然違います。ぜひ試してみてください！',
 NULL, NULL, 89, 15, TRUE, NOW() - INTERVAL '5 days', NOW()),

-- Sakura Ito posts
('seed_post_0000000000000004', 'seed_staff_000000000000002',
 '本日のスタイル✂️ オリーブアッシュのハイライトカラー＋レイヤーカット。光の当たり方で表情が変わるデザインがポイントです！',
 'https://picsum.photos/seed/post_hair_a/600/600', 'image', 124, 23, TRUE, NOW() - INTERVAL '2 days', NOW()),
('seed_post_0000000000000005', 'seed_staff_000000000000002',
 '3月のご予約、残りわずかです！カット＋カラーをご希望の方はお早めにどうぞ 🌷 新宿サロンにてお待ちしています。',
 'https://picsum.photos/seed/post_hair_b/600/600', 'image', 78, 9, TRUE, NOW() - INTERVAL '4 hours', NOW()),
('seed_post_0000000000000006', 'seed_staff_000000000000002',
 '春の新メニュー！イルミナカラー×バレイヤージュがついに解禁🌟 ご希望の方はDMまたは予約ページからどうぞ。',
 'https://picsum.photos/seed/post_hair_c/600/600', 'image', 201, 34, TRUE, NOW() - INTERVAL '6 days', NOW()),

-- Miku Watanabe posts
('seed_post_0000000000000007', 'seed_staff_000000000000003',
 '大阪のお客様、ありがとうございました🌺 今日はニュアンスフレンチ＋押し花アート。大人かわいいデザインに仕上がりました！',
 'https://picsum.photos/seed/post_nail2_a/600/600', 'image', 56, 7, TRUE, NOW() - INTERVAL '1 day', NOW()),
('seed_post_0000000000000008', 'seed_staff_000000000000003',
 '3Dネイルの新作💎 ミラーパウダーとクリスタルストーンの組み合わせ。このデザインは現在受付中です！',
 'https://picsum.photos/seed/post_nail2_b/600/600', 'image', 143, 28, TRUE, NOW() - INTERVAL '3 days', NOW()),

-- Yuna Kimura posts
('seed_post_0000000000000009', 'seed_staff_000000000000004',
 '春は体の緊張がほぐれやすい季節🌿 この機会にしっかりと全身マッサージでリセットしませんか？横浜エリアで出張施術も承ります。',
 'https://picsum.photos/seed/post_spa_a/600/600', 'image', 88, 11, TRUE, NOW() - INTERVAL '2 days', NOW()),
('seed_post_0000000000000010', 'seed_staff_000000000000004',
 'お客様の声 📝「肩こりが嘘みたいに楽になりました」「また来ます！」といただきました。丁寧な施術を心がけています。ご予約お待ちしております🙏',
 NULL, NULL, 112, 19, TRUE, NOW() - INTERVAL '5 days', NOW()),

-- Rina Hayashi posts
('seed_post_0000000000000011', 'seed_staff_000000000000005',
 '本日のブライダルメイク💍 お客様の雰囲気に合わせたピュアウェディングスタイルに仕上げました。素敵な一日になりますように！',
 'https://picsum.photos/seed/post_makeup_a/600/600', 'image', 195, 31, TRUE, NOW() - INTERVAL '1 day', NOW()),
('seed_post_0000000000000012', 'seed_staff_000000000000005',
 '春のトレンドメイク解説✨ 今季のポイントはリップの質感と眉の形。詳しくはメイクレッスンで！60分6,000円からです。',
 'https://picsum.photos/seed/post_makeup_b/600/600', 'image', 267, 45, TRUE, NOW() - INTERVAL '4 days', NOW()),

-- Haruto Yoshida posts
('seed_post_0000000000000013', 'seed_staff_000000000000006',
 '本日のフェードカット✂️ 頭の形を活かしたミッドフェード。スタイリングはマットクレイでタイトに仕上げました。',
 'https://picsum.photos/seed/post_barber_a/600/600', 'image', 73, 14, TRUE, NOW() - INTERVAL '2 days', NOW()),
('seed_post_0000000000000014', 'seed_staff_000000000000006',
 'シェービングの魅力 💈 バーバーシェービングはただ剃るだけじゃない。肌の状態を整え、血行促進にもなります。ぜひ一度体験してみてください。',
 NULL, NULL, 91, 22, TRUE, NOW() - INTERVAL '7 days', NOW()),

-- Aoi Nakamura posts
('seed_post_0000000000000015', 'seed_staff_000000000000007',
 '本日のまつ毛エクステ💫 まつ毛が少ない方でもボリュームが出るようにデザインしました。目元の印象が変わると顔全体が明るく見えます！',
 'https://picsum.photos/seed/post_lash_a/600/600', 'image', 312, 58, TRUE, NOW() - INTERVAL '1 day', NOW()),
('seed_post_0000000000000016', 'seed_staff_000000000000007',
 'まつ毛パーマ before→after 🌟 自まつ毛を使ってここまで変われます！ノーマスカラでこの仕上がり。ぜひプロフィールのギャラリーもご覧ください。',
 'https://picsum.photos/seed/post_lash_b/600/600', 'image', 445, 67, TRUE, NOW() - INTERVAL '3 days', NOW()),
('seed_post_0000000000000017', 'seed_staff_000000000000007',
 '目黒サロン、4月の先行予約を開始しました🌸 人気メニューは早めに埋まります。ご希望の方はお早めにどうぞ！',
 NULL, NULL, 178, 29, TRUE, NOW() - INTERVAL '6 hours', NOW())
ON CONFLICT (id) DO NOTHING;

-- ============================================================
-- FOLLOWS
-- ============================================================

INSERT INTO follows (id, follower_id, followed_id, created_at) VALUES
('seed_flwx_0000000000000001', 'demo_user_0000000000000001', 'demo_staff_000000000000001', NOW() - INTERVAL '25 days'),
('seed_flwx_0000000000000002', 'demo_user_0000000000000001', 'seed_staff_000000000000002', NOW() - INTERVAL '20 days'),
('seed_flwx_0000000000000003', 'demo_user_0000000000000001', 'seed_staff_000000000000007', NOW() - INTERVAL '15 days'),
('seed_flwx_0000000000000004', 'seed_user_0000000000000001', 'demo_staff_000000000000001', NOW() - INTERVAL '22 days'),
('seed_flwx_0000000000000005', 'seed_user_0000000000000001', 'seed_staff_000000000000002', NOW() - INTERVAL '18 days'),
('seed_flwx_0000000000000006', 'seed_user_0000000000000001', 'seed_staff_000000000000003', NOW() - INTERVAL '12 days'),
('seed_flwx_0000000000000007', 'seed_user_0000000000000001', 'seed_staff_000000000000007', NOW() - INTERVAL '8 days'),
('seed_flwx_0000000000000008', 'seed_user_0000000000000002', 'seed_staff_000000000000004', NOW() - INTERVAL '20 days'),
('seed_flwx_0000000000000009', 'seed_user_0000000000000002', 'seed_staff_000000000000007', NOW() - INTERVAL '14 days'),
('seed_flwx_0000000000000010', 'seed_user_0000000000000002', 'demo_staff_000000000000001', NOW() - INTERVAL '10 days'),
('seed_flwx_0000000000000011', 'seed_user_0000000000000003', 'seed_staff_000000000000007', NOW() - INTERVAL '18 days'),
('seed_flwx_0000000000000012', 'seed_user_0000000000000003', 'seed_staff_000000000000006', NOW() - INTERVAL '12 days'),
('seed_flwx_0000000000000013', 'seed_user_0000000000000003', 'seed_staff_000000000000002', NOW() - INTERVAL '6 days'),
('seed_flwx_0000000000000014', 'seed_user_0000000000000004', 'seed_staff_000000000000005', NOW() - INTERVAL '15 days'),
('seed_flwx_0000000000000015', 'seed_user_0000000000000004', 'seed_staff_000000000000007', NOW() - INTERVAL '9 days')
ON CONFLICT (id) DO NOTHING;

-- ============================================================
-- LIKES
-- ============================================================

INSERT INTO likes (id, user_id, post_id, created_at) VALUES
('seed_like_0000000000000001', 'demo_user_0000000000000001', 'seed_post_0000000000000004', NOW() - INTERVAL '2 days'),
('seed_like_0000000000000002', 'demo_user_0000000000000001', 'seed_post_0000000000000015', NOW() - INTERVAL '1 day'),
('seed_like_0000000000000003', 'demo_user_0000000000000001', 'seed_post_0000000000000016', NOW() - INTERVAL '3 days'),
('seed_like_0000000000000004', 'seed_user_0000000000000001', 'seed_post_0000000000000001', NOW() - INTERVAL '3 days'),
('seed_like_0000000000000005', 'seed_user_0000000000000001', 'seed_post_0000000000000004', NOW() - INTERVAL '2 days'),
('seed_like_0000000000000006', 'seed_user_0000000000000001', 'seed_post_0000000000000006', NOW() - INTERVAL '6 days'),
('seed_like_0000000000000007', 'seed_user_0000000000000001', 'seed_post_0000000000000015', NOW() - INTERVAL '1 day'),
('seed_like_0000000000000008', 'seed_user_0000000000000002', 'seed_post_0000000000000009', NOW() - INTERVAL '2 days'),
('seed_like_0000000000000009', 'seed_user_0000000000000002', 'seed_post_0000000000000010', NOW() - INTERVAL '5 days'),
('seed_like_0000000000000010', 'seed_user_0000000000000002', 'seed_post_0000000000000015', NOW() - INTERVAL '1 day'),
('seed_like_0000000000000011', 'seed_user_0000000000000003', 'seed_post_0000000000000015', NOW() - INTERVAL '1 day'),
('seed_like_0000000000000012', 'seed_user_0000000000000003', 'seed_post_0000000000000016', NOW() - INTERVAL '3 days'),
('seed_like_0000000000000013', 'seed_user_0000000000000003', 'seed_post_0000000000000012', NOW() - INTERVAL '4 days'),
('seed_like_0000000000000014', 'seed_user_0000000000000004', 'seed_post_0000000000000013', NOW() - INTERVAL '2 days'),
('seed_like_0000000000000015', 'seed_user_0000000000000004', 'seed_post_0000000000000011', NOW() - INTERVAL '1 day')
ON CONFLICT (id) DO NOTHING;

-- ============================================================
-- COMMENTS
-- ============================================================

INSERT INTO comments (id, post_id, author_id, content, is_active, created_at, updated_at) VALUES
('seed_cmnt_0000000000000001', 'seed_post_0000000000000001', 'seed_user_0000000000000001',
 'めちゃくちゃ可愛い！！次回これにしてもらいたいです😍', TRUE, NOW() - INTERVAL '2 days', NOW()),
('seed_cmnt_0000000000000002', 'seed_post_0000000000000002', 'seed_user_0000000000000002',
 '予約しました！楽しみです✨', TRUE, NOW() - INTERVAL '20 hours', NOW()),
('seed_cmnt_0000000000000003', 'seed_post_0000000000000004', 'demo_user_0000000000000001',
 'アッシュ系カラー素敵すぎる💕 来月予約したいです', TRUE, NOW() - INTERVAL '2 days', NOW()),
('seed_cmnt_0000000000000004', 'seed_post_0000000000000006', 'seed_user_0000000000000001',
 'バレイヤージュずっと気になってました！DM送ります', TRUE, NOW() - INTERVAL '5 days', NOW()),
('seed_cmnt_0000000000000005', 'seed_post_0000000000000015', 'seed_user_0000000000000001',
 'この仕上がりすごいです！来月葵さんに予約入れました🙏', TRUE, NOW() - INTERVAL '1 day', NOW()),
('seed_cmnt_0000000000000006', 'seed_post_0000000000000016', 'seed_user_0000000000000002',
 'before→afterの差がすごい！まつ毛パーマ気になってました', TRUE, NOW() - INTERVAL '3 days', NOW()),
('seed_cmnt_0000000000000007', 'seed_post_0000000000000012', 'seed_user_0000000000000004',
 'メイクレッスン受けてみたいです！詳しくDMしてもいいですか？', TRUE, NOW() - INTERVAL '4 days', NOW()),
('seed_cmnt_0000000000000008', 'seed_post_0000000000000009', 'seed_user_0000000000000002',
 '出張施術できるんですね！横浜ですか？問い合わせたいです', TRUE, NOW() - INTERVAL '2 days', NOW()),
('seed_cmnt_0000000000000009', 'seed_post_0000000000000013', 'seed_user_0000000000000004',
 '仕上がりかっこいい！彼氏に予約させます笑', TRUE, NOW() - INTERVAL '2 days', NOW()),
('seed_cmnt_0000000000000010', 'seed_post_0000000000000011', 'demo_user_0000000000000001',
 '素敵なメイクですね✨ ブライダルで予約可能か聞いてみます', TRUE, NOW() - INTERVAL '1 day', NOW())
ON CONFLICT (id) DO NOTHING;

-- ============================================================
-- TIPS
-- ============================================================

INSERT INTO tips (id, sender_id, recipient_id, staff_profile_id, amount, message, created_at) VALUES
('seed_tip_00000000000000001', 'demo_user_0000000000000001', 'demo_staff_000000000000001', 'demo_staff_prof_000000001',
 500, 'いつもありがとうございます！これからもよろしく💕', NOW() - INTERVAL '19 days'),
('seed_tip_00000000000000002', 'seed_user_0000000000000001', 'seed_staff_000000000000002', 'seed_prof_0000000000000001',
 1000, 'カラーがすごく気に入りました！ありがとうございます🌸', NOW() - INTERVAL '17 days'),
('seed_tip_00000000000000003', 'seed_user_0000000000000002', 'seed_staff_000000000000004', 'seed_prof_0000000000000003',
 2000, '肩こりが本当に楽になりました。感謝です！', NOW() - INTERVAL '13 days'),
('seed_tip_00000000000000004', 'seed_user_0000000000000003', 'seed_staff_000000000000007', 'seed_prof_0000000000000006',
 3000, 'いつも丁寧な施術ありがとうございます✨', NOW() - INTERVAL '9 days'),
('seed_tip_00000000000000005', 'seed_user_0000000000000004', 'seed_staff_000000000000006', 'seed_prof_0000000000000005',
 500, 'バーバー最高でした！また来ます', NOW() - INTERVAL '7 days'),
('seed_tip_00000000000000006', 'demo_user_0000000000000001', 'seed_staff_000000000000007', 'seed_prof_0000000000000006',
 1500, '投稿いつも参考にしてます！応援しています', NOW() - INTERVAL '4 days'),
('seed_tip_00000000000000007', 'seed_user_0000000000000001', 'seed_staff_000000000000007', 'seed_prof_0000000000000006',
 2000, 'まつ毛パーマ予約しました！楽しみです', NOW() - INTERVAL '2 days'),
('seed_tip_00000000000000008', 'seed_user_0000000000000003', 'seed_staff_000000000000002', 'seed_prof_0000000000000001',
 1000, 'SNSの投稿がいつも素敵です！応援してます', NOW() - INTERVAL '1 day')
ON CONFLICT (id) DO NOTHING;

-- ============================================================
-- POINT TRANSACTIONS (matching tips + some purchases)
-- ============================================================

INSERT INTO point_transactions (id, user_id, type, amount, balance_after, reference_id, created_at) VALUES
-- demo_user tip sends
('seed_ptxn_0000000000000001', 'demo_user_0000000000000001', 'tip_sent', -500, 9500, 'seed_tip_00000000000000001', NOW() - INTERVAL '19 days'),
('seed_ptxn_0000000000000002', 'demo_user_0000000000000001', 'tip_sent', -1500, 8000, 'seed_tip_00000000000000006', NOW() - INTERVAL '4 days'),
-- demo_staff tip receives
('seed_ptxn_0000000000000003', 'demo_staff_000000000000001', 'tip_received', 500, 5500, 'seed_tip_00000000000000001', NOW() - INTERVAL '19 days'),
-- Sakura tip receives
('seed_ptxn_0000000000000004', 'seed_staff_000000000000002', 'tip_received', 1000, 8200, 'seed_tip_00000000000000002', NOW() - INTERVAL '17 days'),
('seed_ptxn_0000000000000005', 'seed_staff_000000000000002', 'tip_received', 1000, 9200, 'seed_tip_00000000000000008', NOW() - INTERVAL '1 day'),
-- Yuna tip receives
('seed_ptxn_0000000000000006', 'seed_staff_000000000000004', 'tip_received', 2000, 9100, 'seed_tip_00000000000000003', NOW() - INTERVAL '13 days'),
-- Aoi tip receives
('seed_ptxn_0000000000000007', 'seed_staff_000000000000007', 'tip_received', 3000, 11500, 'seed_tip_00000000000000004', NOW() - INTERVAL '9 days'),
('seed_ptxn_0000000000000008', 'seed_staff_000000000000007', 'tip_received', 1500, 13000, 'seed_tip_00000000000000006', NOW() - INTERVAL '4 days'),
('seed_ptxn_0000000000000009', 'seed_staff_000000000000007', 'tip_received', 2000, 15000, 'seed_tip_00000000000000007', NOW() - INTERVAL '2 days'),
-- Haruto tip receives
('seed_ptxn_0000000000000010', 'seed_staff_000000000000006', 'tip_received', 500, 3800, 'seed_tip_00000000000000005', NOW() - INTERVAL '7 days'),
-- User point purchases
('seed_ptxn_0000000000000011', 'seed_user_0000000000000001', 'purchase', 2000, 2000, NULL, NOW() - INTERVAL '22 days'),
('seed_ptxn_0000000000000012', 'seed_user_0000000000000001', 'tip_sent', -1000, 1000, 'seed_tip_00000000000000002', NOW() - INTERVAL '17 days'),
('seed_ptxn_0000000000000013', 'seed_user_0000000000000001', 'tip_sent', -2000, -1000, 'seed_tip_00000000000000007', NOW() - INTERVAL '2 days'),
('seed_ptxn_0000000000000014', 'seed_user_0000000000000002', 'purchase', 3000, 3000, NULL, NOW() - INTERVAL '15 days'),
('seed_ptxn_0000000000000015', 'seed_user_0000000000000002', 'tip_sent', -2000, 1000, 'seed_tip_00000000000000003', NOW() - INTERVAL '13 days'),
('seed_ptxn_0000000000000016', 'seed_user_0000000000000003', 'purchase', 5000, 5000, NULL, NOW() - INTERVAL '12 days'),
('seed_ptxn_0000000000000017', 'seed_user_0000000000000003', 'tip_sent', -3000, 2000, 'seed_tip_00000000000000004', NOW() - INTERVAL '9 days'),
('seed_ptxn_0000000000000018', 'seed_user_0000000000000003', 'tip_sent', -1000, 1000, 'seed_tip_00000000000000008', NOW() - INTERVAL '1 day'),
('seed_ptxn_0000000000000019', 'seed_user_0000000000000004', 'purchase', 1000, 1000, NULL, NOW() - INTERVAL '10 days'),
('seed_ptxn_0000000000000020', 'seed_user_0000000000000004', 'tip_sent', -500, 500, 'seed_tip_00000000000000005', NOW() - INTERVAL '7 days')
ON CONFLICT (id) DO NOTHING;

-- ============================================================
-- UPDATE followers_count on staff_profiles to match seed follows
-- ============================================================

UPDATE staff_profiles SET followers_count = (
    SELECT COUNT(*) FROM follows WHERE followed_id = (
        SELECT user_id FROM staff_profiles sp2 WHERE sp2.id = staff_profiles.id
    )
)
WHERE id IN (
    'demo_staff_prof_000000001',
    'seed_prof_0000000000000001',
    'seed_prof_0000000000000002',
    'seed_prof_0000000000000003',
    'seed_prof_0000000000000004',
    'seed_prof_0000000000000005',
    'seed_prof_0000000000000006'
);
