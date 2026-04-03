-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- +goose Down
DROP TABLE IF EXISTS users;

-- Example: create a users table (will be extended later)
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT uni_users_email UNIQUE (email)
);

INSERT INTO users (id, email, password_hash, created_at, updated_at)
VALUES (
           uuid_generate_v4(),
           'admin@pesamind.app',
           '$2a$10$eBHSyDfqbdm7hSOBcGUr4u1IeihTlH5yF49VOdIKu10CWNeH8nESe', -- bcrypt hash for 'admin1234'
           NOW(),
           NOW()
       );

