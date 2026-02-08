CREATE TABLE IF NOT EXISTS amilate.appointments (
    id SERIAL PRIMARY KEY,
    public_id UUID UNIQUE NOT NULL,
    user_id INT NOT NULL REFERENCES amilate.users(id) ON DELETE CASCADE,
    destination_id INT NOT NULL REFERENCES amilate.destinations(id) ON DELETE CASCADE,

    scheduled_at VARCHAR(255) NOT NULL,
    transport_mode VARCHAR(50) NOT NULL CHECK (transport_mode IN ('car', 'public_transport', 'foot')),

    estimated_travel_minutes BIGINT NOT NULL,
    estimated_travel_range BIGINT NOT NULL,
    last_travel_update_at VARCHAR(255) NOT NULL,

    notification_state VARCHAR(50) NOT NULL CHECK (notification_state IN ('pending', 'monitoring', 'sent', 'cancelled')),
    status VARCHAR(50) NOT NULL CHECK (status IN ('scheduled', 'completed', 'cancelled')),

    delete_after VARCHAR(255) NOT NULL,
    geo_fence_id VARCHAR(255) NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE INDEX idx_appointments_user_id ON amilate.appointments(user_id);
CREATE INDEX idx_appointments_destination_id ON amilate.appointments(destination_id);
CREATE INDEX idx_appointments_public_id ON amilate.appointments(public_id);
CREATE INDEX idx_appointments_status ON amilate.appointments(status);
CREATE INDEX idx_appointments_notification_state ON amilate.appointments(notification_state);
CREATE INDEX idx_appointments_scheduled_at ON amilate.appointments(scheduled_at);