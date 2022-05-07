package service_impl

import "gitlab.ozon.dev/zBlur/homework-2/internal/service"

type Service struct {
	userService *UserService
}

func New() *Service {
	return &Service{}
}

func (s *Service) User() service.UserService {
	if s.userService != nil {
		return s.userService
	}

	s.userService = &UserService{}

	return s.userService
}
