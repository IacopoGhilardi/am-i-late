CREATE TABLE IF NOT EXISTS amilate.users (
    id SERIAL PRIMARY KEY,
    public_id UUID UNIQUE NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    age_confirmed BOOLEAN NOT NULL DEFAULT FALSE,
    privacy_accepted BOOLEAN NOT NULL DEFAULT FALSE,
    terms_accepted BOOLEAN NOT NULL DEFAULT FALSE,
    location_permission BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE INDEX idx_users_email ON amilate.users(email);
CREATE INDEX idx_users_public_id ON amilate.users(public_id);