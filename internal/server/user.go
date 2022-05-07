package server

import (
	"context"
	"errors"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	pb "gitlab.ozon.dev/zBlur/homework-2/pkg/api"
)

var (
	ErrorTimeOut = errors.New("timeout error")
)

func (t tserver) Retrieve(ctx context.Context, req *pb.RetrieveUserRequest) (*pb.User, error) {
	userId := domain.UserId(req.GetId())
	channel := t.serv.User().Retrieve(userId, t.repo.User())

	select {
	case userRetrieved := <-channel:
		if userRetrieved.Error != nil {
			return nil, userRetrieved.Error
		}
		user := pb.User{
			Id:        userRetrieved.User.Id.ToInt64(),
			UserName:  userRetrieved.User.UserName,
			FirstName: userRetrieved.User.FirstName,
			LastName:  userRetrieved.User.LastName,
		}

		return &user, nil

	case <-ctx.Done():
		return nil, ErrorTimeOut
	}
}

func (t tserver) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.Empty, error) {
	user := domain.User{
		Id:        domain.UserId(req.GetId()),
		UserName:  req.GetUserName(),
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
	}
	channel := t.serv.User().Create(&user, t.repo.User())

	select {
	case err := <-channel:
		if err != nil {
			return nil, err
		}
		return &pb.Empty{}, nil

	case <-ctx.Done():
		return nil, ErrorTimeOut
	}
}

func (t tserver) RetrieveOrCreate(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	user := domain.User{
		Id:        domain.UserId(req.GetId()),
		UserName:  req.GetUserName(),
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
	}
	channel := t.serv.User().RetrieveOrCreate(&user, t.repo.User())

	select {
	case userRetrieved := <-channel:
		if userRetrieved.Error != nil {
			return nil, userRetrieved.Error
		}
		user := pb.User{
			Id:        userRetrieved.User.Id.ToInt64(),
			UserName:  userRetrieved.User.UserName,
			FirstName: userRetrieved.User.FirstName,
			LastName:  userRetrieved.User.LastName,
		}
		return &user, nil

	case <-ctx.Done():
		return nil, ErrorTimeOut
	}
}

func (t tserver) Update(ctx context.Context, req *pb.UpdateUserRequest) (*pb.Empty, error) {
	user := domain.User{
		Id:        domain.UserId(req.GetId()),
		UserName:  req.GetUserName(),
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
	}
	channel := t.serv.User().Update(&user, t.repo.User())

	select {
	case err := <-channel:
		if err != nil {
			return nil, err
		}
		return &pb.Empty{}, nil

	case <-ctx.Done():
		return nil, ErrorTimeOut
	}
}

func (t tserver) Delete(ctx context.Context, req *pb.DeleteUserRequest) (*pb.Empty, error) {
	userId := domain.UserId(req.GetId())
	channel := t.serv.User().Delete(userId, t.repo.User())

	select {
	case err := <-channel:
		if err != nil {
			return nil, err
		}
		return &pb.Empty{}, nil

	case <-ctx.Done():
		return nil, ErrorTimeOut
	}
}
