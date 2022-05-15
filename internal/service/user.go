package service

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"gitlab.ozon.dev/zBlur/homework-2/internal/repository"
)

type UserService interface {
	Create(ctx context.Context, user *domain.User, repo repository.UserRepository) error
	Retrieve(ctx context.Context, userId domain.UserId, repo repository.UserRepository) domain.UserRetrieve
	RetrieveOrCreate(ctx context.Context, user *domain.User, repo repository.UserRepository) domain.UserRetrieve
	Update(ctx context.Context, user *domain.User, repo repository.UserRepository) error
	Delete(ctx context.Context, userId domain.UserId, repo repository.UserRepository) error
}
