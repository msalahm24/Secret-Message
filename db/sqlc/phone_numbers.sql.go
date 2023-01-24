// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: phone_numbers.sql

package db

import (
	"context"
	"database/sql"
)

const createPhoneNumber = `-- name: CreatePhoneNumber :one
INSERT INTO phone_numbers (
    phone_number,
    location,
    user_id
) VALUES (
    $1,
    $2,
    $3
) RETURNING id, phone_number, location, user_id, created_at
`

type CreatePhoneNumberParams struct {
	PhoneNumber string         `json:"phoneNumber"`
	Location    sql.NullString `json:"location"`
	UserID      string         `json:"userID"`
}

func (q *Queries) CreatePhoneNumber(ctx context.Context, arg CreatePhoneNumberParams) (PhoneNumber, error) {
	row := q.db.QueryRowContext(ctx, createPhoneNumber, arg.PhoneNumber, arg.Location, arg.UserID)
	var i PhoneNumber
	err := row.Scan(
		&i.ID,
		&i.PhoneNumber,
		&i.Location,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}

const getPhoneNumbersByUser = `-- name: GetPhoneNumbersByUser :many
SELECT id, phone_number, location, user_id, created_at FROM phone_numbers
WHERE user_id = $1
`

func (q *Queries) GetPhoneNumbersByUser(ctx context.Context, userID string) ([]PhoneNumber, error) {
	rows, err := q.db.QueryContext(ctx, getPhoneNumbersByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []PhoneNumber{}
	for rows.Next() {
		var i PhoneNumber
		if err := rows.Scan(
			&i.ID,
			&i.PhoneNumber,
			&i.Location,
			&i.UserID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByphoneNumber = `-- name: GetUserByphoneNumber :one
SELECT id, phone_number, location, user_id, created_at FROM phone_numbers
WHERE phone_number = $1
`

func (q *Queries) GetUserByphoneNumber(ctx context.Context, phoneNumber string) (PhoneNumber, error) {
	row := q.db.QueryRowContext(ctx, getUserByphoneNumber, phoneNumber)
	var i PhoneNumber
	err := row.Scan(
		&i.ID,
		&i.PhoneNumber,
		&i.Location,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}
