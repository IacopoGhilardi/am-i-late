CREATE TABLE IF NOT EXISTS amilate.destinations (
    id SERIAL PRIMARY KEY,
    public_id UUID UNIQUE NOT NULL,
    user_id INT NOT NULL REFERENCES amilate.users(id) ON DELETE CASCADE,
    full_address VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    formatted_address VARCHAR(255) NOT NULL,
    google_place_id VARCHAR(255) NOT NULL,
    latitude DECIMAL(10,7) NOT NULL,
    longitude DECIMAL(10,7) NOT NULL,
    time_zone VARCHAR(255) NOT NULL,
    is_saved BOOLEAN DEFAULT FALSE,
    last_used_at TIMESTAMPTZ,
    delete_after VARCHAR(255) NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE INDEX idx_destinations_user_id ON amilate.destinations(user_id);
CREATE INDEX idx_destinations_public_id ON amilate.destinations(public_id);
CREATE INDEX idx_destinations_google_place_id ON amilate.destinations(google_place_id);
CREATE INDEX idx_destinations_last_used_at ON amilate.destinations(last_used_at);