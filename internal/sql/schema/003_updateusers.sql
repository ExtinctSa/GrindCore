-- +goose Up
ALTER TABLE users
ADD COLUMN updated_at TIMESTAMP NOT NULL DEFAULT now();

-- +goose Down
ALTER TABLE users
DROP COLUMN updated_at;