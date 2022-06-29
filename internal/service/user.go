package service

import (
	"context"
	"github.com/abrbird/portfolio_bot/internal/domain"
	"github.com/abrbird/portfolio_bot/internal/repository"
)

type UserService interface {
	Create(ctx context.Context, user *domain.User, repo repository.UserRepository) error
	Retrieve(ctx context.Context, userId domain.UserId, repo repository.UserRepository) domain.UserRetrieve
	RetrieveOrCreate(ctx context.Context, user *domain.User, repo repository.UserRepository) domain.UserRetrieve
	Update(ctx context.Context, user *domain.User, repo repository.UserRepository) error
	Delete(ctx context.Context, userId domain.UserId, repo repository.UserRepository) error
}
