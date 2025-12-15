-- name: CreateUser :one
INSERT INTO users (id, username, email, hashed_password, created_at)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;

-- name: CreateUserByEmail :one
INSERT INTO users (email, hashed_password, username)
VALUES ($1, $2, $3)
RETURNING id, username, email, hashed_password, created_at;

-- name: GetUserByUsername :one
SELECT id, created_at, email, hashed_password, username
FROM users
WHERE username = $1;

-- name: GetUserByID :one
SELECT id, created_at, email, username
FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET email = $2,
    hashed_password = $3,
    updated_at = NOW()
WHERE id = $1
RETURNING id, email, username, created_at, updated_at;