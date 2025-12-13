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
INSERT INTO users (email, hashed_password)
VALUES ($1, $2)
RETURNING id, username, email, hashed_password, created_at;

-- name: GetUserByUsername :one
SELECT id, created_at, email, hashed_password, username
FROM users
WHERE username = $1;