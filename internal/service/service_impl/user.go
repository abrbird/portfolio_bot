package service_impl

import (
	"context"
	"github.com/abrbird/portfolio_bot/internal/domain"
	"github.com/abrbird/portfolio_bot/internal/repository"
)

type UserService struct{}

func (us UserService) Create(ctx context.Context, user *domain.User, repo repository.UserRepository) error {
	return repo.Create(ctx, user)
}

func (us UserService) Retrieve(ctx context.Context, userId domain.UserId, repo repository.UserRepository) domain.UserRetrieve {
	return repo.Retrieve(ctx, userId)
}

func (us UserService) RetrieveOrCreate(ctx context.Context, user *domain.User, repo repository.UserRepository) domain.UserRetrieve {
	return repo.RetrieveOrCreate(ctx, user)
}

func (us UserService) Update(ctx context.Context, user *domain.User, repo repository.UserRepository) error {
	return repo.Update(ctx, user)
}

func (us UserService) Delete(ctx context.Context, userId domain.UserId, repo repository.UserRepository) error {
	return repo.Delete(ctx, userId)
}
