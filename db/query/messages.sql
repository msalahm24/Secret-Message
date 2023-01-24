-- name: CreateMessage :one
INSERT INTO messages (
    user_id,
    to_phone_number,
    content
) VALUES (
    $1,
    $2,
    $3
) RETURNING *;


-- name: GetMessageByUser :many
SELECT * FROM messages
WHERE user_id = $1;

-- name: GetMessageByToPhoneNumber :many
SELECT * FROM messages
WHERE to_phone_number = $1;

-- name: CountMessages :one
SELECT count(*) FROM messages;