-- name: CreateHabit :one
INSERT INTO habits (habitName, frequency, category, user_id)
VALUES ($1, $2, $3, $4)
RETURNING id, habitName, frequency, category, created_at, updated_at, user_id;

-- name: GetAllHabits :many
SELECT * FROM habits
ORDER BY created_at ASC;

-- name: DeleteHabit :exec
DELETE FROM habits
WHERE id = $1;

-- name: GetHabitByCategory :many
SELECT * FROM habits
WHERE category = $1
ORDER BY created_at ASC;