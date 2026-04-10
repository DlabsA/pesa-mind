-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT uni_users_email UNIQUE (email)
);

-- Create profiles table (one-to-one with users)
CREATE TABLE IF NOT EXISTS profiles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    username VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL DEFAULT 'Free',
    balance NUMERIC(18,2) NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT fk_profiles_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT uni_profiles_username UNIQUE (username),
    CONSTRAINT uni_profiles_userid UNIQUE (user_id)
);

-- Seed default admin user if not exists
INSERT INTO users (id, email, password_hash, created_at, updated_at)
SELECT uuid_generate_v4(), 'admin@pesamind.app',
       '$2a$10$eBHSyDfqbdm7hSOBcGUr4u1IeihTlH5yF49VOdIKu10CWNeH8nESe', -- bcrypt for 'admin1234'
       NOW(), NOW()
WHERE NOT EXISTS (SELECT 1 FROM users WHERE email = 'admin@pesamind.app');

-- Seed admin profile for admin user if not exists
INSERT INTO profiles (id, user_id, username, type, balance, created_at, updated_at)
SELECT uuid_generate_v4(), u.id, 'admin', 'Free', 0, NOW(), NOW()
FROM users u
WHERE u.email = 'admin@pesamind.app'
  AND NOT EXISTS (SELECT 1 FROM profiles p WHERE p.user_id = u.id);


