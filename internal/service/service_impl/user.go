package service_impl

import (
	"errors"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"gitlab.ozon.dev/zBlur/homework-2/internal/repository"
)

type UserService struct{}

func (t UserService) RetrieveOrCreate(user *domain.User, repo repository.UserRepository) <-chan domain.UserRetrieve {
	channel := make(chan domain.UserRetrieve)

	go func() {
		err := repo.Update(user)
		if err == nil {
			channel <- domain.UserRetrieve{User: user, Error: nil}
		} else if errors.Is(err, domain.ErrorNotFound) {
			err := repo.Create(user)
			if err != nil {
				channel <- domain.UserRetrieve{User: nil, Error: err}
			}
			channel <- domain.UserRetrieve{User: user, Error: nil}
		} else {
			channel <- domain.UserRetrieve{User: nil, Error: err}
		}
		close(channel)
	}()

	return channel
}

func (t UserService) Create(user *domain.User, repo repository.UserRepository) <-chan error {
	channel := make(chan error)

	go func() {
		channel <- repo.Create(user)
		close(channel)
	}()

	return channel
}

func (t UserService) Retrieve(userId domain.UserId, repo repository.UserRepository) <-chan domain.UserRetrieve {
	channel := make(chan domain.UserRetrieve)

	go func() {
		channel <- repo.Retrieve(userId)
		close(channel)
	}()

	return channel
}

func (t UserService) Update(user *domain.User, repo repository.UserRepository) <-chan error {
	channel := make(chan error)

	go func() {
		channel <- repo.Update(user)
		close(channel)
	}()

	return channel
}

func (t UserService) Delete(userId domain.UserId, repo repository.UserRepository) <-chan error {
	channel := make(chan error)

	go func() {
		channel <- repo.Delete(userId)
		close(channel)
	}()

	return channel
}
