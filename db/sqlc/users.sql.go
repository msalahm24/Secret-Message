// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: users.sql

package db

import (
	"context"
	"database/sql"
)

const countUsers = `-- name: CountUsers :one
SELECT count(*) FROM users
`

func (q *Queries) CountUsers(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countUsers)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createUser = `-- name: CreateUser :one
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
) RETURNING id, user_name, phone_number, location, type_of_device, created_at
`

type CreateUserParams struct {
	UserName     string         `json:"userName"`
	PhoneNumber  string         `json:"phoneNumber"`
	Location     sql.NullString `json:"location"`
	TypeOfDevice sql.NullString `json:"typeOfDevice"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.UserName,
		arg.PhoneNumber,
		arg.Location,
		arg.TypeOfDevice,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserName,
		&i.PhoneNumber,
		&i.Location,
		&i.TypeOfDevice,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, user_name, phone_number, location, type_of_device, created_at FROM users
WHERE user_name = $1
`

func (q *Queries) GetUser(ctx context.Context, userName string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, userName)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserName,
		&i.PhoneNumber,
		&i.Location,
		&i.TypeOfDevice,
		&i.CreatedAt,
	)
	return i, err
}
