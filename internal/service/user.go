package service

import (
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"gitlab.ozon.dev/zBlur/homework-2/internal/repository"
)

type UserService interface {
	Create(user *domain.User, repo repository.UserRepository) <-chan error
	Retrieve(userId domain.UserId, repo repository.UserRepository) <-chan domain.UserRetrieve
	RetrieveOrCreate(user *domain.User, repo repository.UserRepository) <-chan domain.UserRetrieve
	Update(user *domain.User, repo repository.UserRepository) <-chan error
	Delete(userId domain.UserId, repo repository.UserRepository) <-chan error
}
