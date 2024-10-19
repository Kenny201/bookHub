package service

import (
	"context"

	"user_service/internals/models"
	"user_service/internals/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
	return s.repo.Create(ctx, user)
}

func (s *UserService) FindUser(ctx context.Context, id int) (*models.User, error) {
	return s.repo.GetByID(ctx, id)
}
