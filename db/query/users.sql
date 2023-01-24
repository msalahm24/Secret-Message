-- name: CreateUser :one
INSERT INTO users (
    user_name,
    phone_number,
    location,
    type_of_device
) VALUES (
    $1,
    $2,
    $3,
    $4
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE user_name = $1;

-- name: CountUsers :one
SELECT count(*) FROM users;