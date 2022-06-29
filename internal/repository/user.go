package repository

import (
	"context"
	"github.com/abrbird/portfolio_bot/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	Retrieve(ctx context.Context, userId domain.UserId) domain.UserRetrieve
	RetrieveOrCreate(ctx context.Context, user *domain.User) domain.UserRetrieve
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, userId domain.UserId) error
}
