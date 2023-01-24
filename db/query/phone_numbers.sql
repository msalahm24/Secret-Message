-- name: CreatePhoneNumber :one
INSERT INTO phone_numbers (
    phone_number,
    location,
    user_id
) VALUES (
    $1,
    $2,
    $3
) RETURNING *;

-- name: GetPhoneNumbersByUser :many
SELECT * FROM phone_numbers
WHERE user_id = $1;


-- name: GetUserByphoneNumber :one
SELECT * FROM phone_numbers
WHERE phone_number = $1;

-- name: CountMessages :one
SELECT count(*) FROM phone_numbers;