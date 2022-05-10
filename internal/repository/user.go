package repository

import (
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
)

type UserRepository interface {
	Create(user *domain.User) error
	Retrieve(userId domain.UserId) domain.UserRetrieve
	RetrieveOrCreate(user *domain.User) domain.UserRetrieve
	Update(user *domain.User) error
	Delete(userId domain.UserId) error
}
