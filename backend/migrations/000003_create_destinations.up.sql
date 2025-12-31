CREATE TABLE IF NOT EXISTS amilate.destinations (
    id SERIAL PRIMARY KEY,
    public_id UUID UNIQUE NOT NULL,
    user_id INT NOT NULL REFERENCES amilate.users(id) ON DELETE CASCADE,
    full_address VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    postal_code VARCHAR(20),
    country VARCHAR(100) NOT NULL,
    latitude DECIMAL(9,6) NOT NULL,
    longitude DECIMAL(9,6) NOT NULL,
    transport_mode VARCHAR(255) NOT NULL,
    time_zone VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);
