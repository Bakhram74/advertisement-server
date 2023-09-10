package repository

import (
	"context"
	db "github.com/Bakhram74/advertisement-server.git/db/sqlc"
)

type AuthRepository struct {
	store Store
}

func NewAuthRepository(store Store) *AuthRepository {
	return &AuthRepository{
		store: store,
	}
}

func (a AuthRepository) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	user, err := a.store.CreateUser(ctx, arg)
	if err != nil {
		return db.User{}, err
	}
	return user, err
}

func (a AuthRepository) GetUser(ctx context.Context, phoneNumber string) (db.User, error) {
	user, err := a.store.GetUser(ctx, phoneNumber)
	if err != nil {
		return db.User{}, err
	}
	return user, err
}
