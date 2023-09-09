package service

import (
	"context"
	db "github.com/Bakhram74/advertisement-server.git/db/sqlc"
	"github.com/Bakhram74/advertisement-server.git/internal/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (service *AuthService) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	return service.repo.CreateUser(ctx, arg)
}
