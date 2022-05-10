package service_impl

import (
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"gitlab.ozon.dev/zBlur/homework-2/internal/repository"
)

type UserService struct{}

//func (us UserService) Create(user *domain.User, repo repository.UserRepository) <-chan error {
//	channel := make(chan error)
//
//	go func() {
//		channel <- repo.Create(user)
//		close(channel)
//	}()
//
//	return channel
//}
//
//func (us UserService) Retrieve(userId domain.UserId, repo repository.UserRepository) <-chan domain.UserRetrieve {
//	channel := make(chan domain.UserRetrieve)
//
//	go func() {
//		channel <- repo.Retrieve(userId)
//		close(channel)
//	}()
//
//	return channel
//}
//
//func (us UserService) RetrieveOrCreate(user *domain.User, repo repository.UserRepository) <-chan domain.UserRetrieve {
//	channel := make(chan domain.UserRetrieve)
//
//	go func() {
//		channel <- repo.RetrieveOrCreate(user)
//		close(channel)
//	}()
//
//	return channel
//}
//
//func (us UserService) Update(user *domain.User, repo repository.UserRepository) <-chan error {
//	channel := make(chan error)
//
//	go func() {
//		channel <- repo.Update(user)
//		close(channel)
//	}()
//
//	return channel
//}
//
//func (us UserService) Delete(userId domain.UserId, repo repository.UserRepository) <-chan error {
//	channel := make(chan error)
//
//	go func() {
//		channel <- repo.Delete(userId)
//		close(channel)
//	}()
//
//	return channel
//}

func (us UserService) Create(user *domain.User, repo repository.UserRepository) error {
	return repo.Create(user)
}

func (us UserService) Retrieve(userId domain.UserId, repo repository.UserRepository) domain.UserRetrieve {
	return repo.Retrieve(userId)
}

func (us UserService) RetrieveOrCreate(user *domain.User, repo repository.UserRepository) domain.UserRetrieve {
	return repo.RetrieveOrCreate(user)
}

func (us UserService) Update(user *domain.User, repo repository.UserRepository) error {
	return repo.Update(user)
}

func (us UserService) Delete(userId domain.UserId, repo repository.UserRepository) error {
	return repo.Delete(userId)
}
