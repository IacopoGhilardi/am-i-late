CREATE TABLE IF NOT EXISTS amilate.alarms (
    id SERIAL PRIMARY KEY,
    time VARCHAR(10) NOT NULL,
    user_id INT NOT NULL REFERENCES amilate.users(id) ON DELETE CASCADE,
    destination_id INT REFERENCES amilate.destinations(id) ON DELETE SET NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);