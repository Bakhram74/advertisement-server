package repository

import (
	"context"
	db "github.com/Bakhram74/advertisement-server.git/db/sqlc"
)

type Authorization interface {
	CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error)
	GetUser(ctx context.Context, phoneNumber string) (db.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(store Store) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(store),
	}
}
