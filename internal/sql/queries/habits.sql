-- name: CreateHabit :one
INSERT INTO habits (habitName, frequency, category, user_id)
VALUES ($1, $2, $3, $4)
RETURNING id, habitName, frequency, category, created_at, updated_at, user_id;

-- name: GetAllHabits :many
SELECT * 
FROM habits
WHERE user_id = $1
ORDER BY created_at ASC;

-- name: DeleteHabit :exec
DELETE FROM habits
WHERE user_id = $1
    AND id = $2;

-- name: GetHabitByCategory :many
SELECT * 
FROM habits
WHERE user_id = $1 
    AND category = $2
ORDER BY created_at ASC;

-- name: ListHabitsWithoutCategory :many
SELECT *
FROM habits
WHERE user_id = $1
    AND category IS NULL
ORDER BY created_at ASC;

-- name: GetHabitByID :one
SELECT *
FROM habits
WHERE id = $1;

-- name: UpdateHabit :one
UPDATE habits
SET
    habitName = COALESCE(sqlc.narg(habitName), habitName),
    frequency = COALESCE(sqlc.narg(frequency), frequency),
    category  = COALESCE(sqlc.narg(category), category),
    updated_at = NOW()
WHERE id = $1
  AND user_id = $2
RETURNING id, habitName, frequency, category, created_at, updated_at, user_id;
