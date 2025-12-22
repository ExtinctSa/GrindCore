-- name: MarkCompletion :exec
INSERT INTO completion (habit_id, user_id, completed_date)
VALUES ($1, $2, CURRENT_DATE)
ON CONFLICT (habit_id, completed_date) DO NOTHING;

-- name: CheckCompletion :one
SELECT EXISTS (
    SELECT 1
    FROM completion
    WHERE habit_id = $1
        AND user_id = $2
        AND completed_date = CURRENT_DATE
);
