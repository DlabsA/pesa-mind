-- +goose Up
-- Add an initial user for development/testing
INSERT INTO users (id, email, password_hash, created_at, updated_at)
VALUES (
    uuid_generate_v4(),
    'admin@pesamind.app',
    '$2a$10$nKfXxw/QZ3O9II0U8Gx6tOv1LyD/8w80GGdqDOqShoe7bX7lLF07G', -- bcrypt hash for 'admin12345'
    NOW(),
    NOW()
);

-- +goose Down
DELETE FROM users WHERE email = 'admin@pesamind.app';

