// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID          int64
	Username    string
	PhoneNumber string
	Password    string
	Role        string
	IsBanned    bool
	CreatedAt   pgtype.Timestamptz
}
