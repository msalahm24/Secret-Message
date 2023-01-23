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