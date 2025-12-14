-- Services table
CREATE TABLE services (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    price NUMERIC(20,2) DEFAULT 0.0,
    is_active BOOLEAN NOT NULL DEFAULT false,
    created_by UUID NULL REFERENCES users(id) ON DELETE SET NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);


-- Create an index on the username column
CREATE INDEX idx_services_name ON services(name);