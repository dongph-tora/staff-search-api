CREATE TABLE IF NOT EXISTS users (
    id                      VARCHAR(26) PRIMARY KEY,
    email                   VARCHAR(255) NOT NULL UNIQUE,
    password_hash           VARCHAR(255),
    name                    VARCHAR(255) NOT NULL,
    phone_number            VARCHAR(20),
    avatar_url              TEXT,
    bio                     TEXT,
    role                    VARCHAR(20) NOT NULL DEFAULT 'user',
    is_staff                BOOLEAN NOT NULL DEFAULT FALSE,
    is_staff_registered     BOOLEAN NOT NULL DEFAULT FALSE,
    is_verified             BOOLEAN NOT NULL DEFAULT FALSE,
    google_id               VARCHAR(255) UNIQUE,
    apple_id                VARCHAR(255) UNIQUE,
    auth_provider           VARCHAR(50) NOT NULL DEFAULT 'email',
    status                  VARCHAR(20) NOT NULL DEFAULT 'active',
    points                  INTEGER NOT NULL DEFAULT 0,
    privacy_policy_accepted BOOLEAN NOT NULL DEFAULT FALSE,
    privacy_policy_version  VARCHAR(20),
    last_login_at           TIMESTAMPTZ,
    created_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_users_email ON users (email);
CREATE INDEX IF NOT EXISTS idx_users_role ON users (role);
CREATE INDEX IF NOT EXISTS idx_users_status ON users (status);
CREATE INDEX IF NOT EXISTS idx_users_google_id ON users (google_id) WHERE google_id IS NOT NULL;
CREATE INDEX IF NOT EXISTS idx_users_apple_id ON users (apple_id) WHERE apple_id IS NOT NULL;
