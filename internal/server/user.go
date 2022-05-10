package server

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	pb "gitlab.ozon.dev/zBlur/homework-2/pkg/api"
)

//func (t tserver) RetrieveUser(ctx context.Context, req *pb.RetrieveUserRequest) (*pb.User, error) {
//	userId := domain.UserId(req.GetId())
//	channel := t.serv.User().Retrieve(userId, t.repo.User())
//
//	select {
//	case userRetrieved := <-channel:
//		if userRetrieved.Error != nil {
//			return nil, userRetrieved.Error
//		}
//		user := pb.User{
//			Id:        userRetrieved.User.Id.ToInt64(),
//			UserName:  userRetrieved.User.UserName,
//			FirstName: userRetrieved.User.FirstName,
//			LastName:  userRetrieved.User.LastName,
//		}
//
//		return &user, nil
//
//	case <-ctx.Done():
//		return nil, ErrorTimeOut
//	}
//}
//
//func (t tserver) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.Empty, error) {
//	user := domain.User{
//		Id:        domain.UserId(req.GetId()),
//		UserName:  req.GetUserName(),
//		FirstName: req.GetFirstName(),
//		LastName:  req.GetLastName(),
//	}
//	channel := t.serv.User().Create(&user, t.repo.User())
//
//	select {
//	case err := <-channel:
//		if err != nil {
//			return nil, err
//		}
//		return &pb.Empty{}, nil
//
//	case <-ctx.Done():
//		return nil, ErrorTimeOut
//	}
//}
//
//func (t tserver) RetrieveOrCreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
//	user := domain.User{
//		Id:        domain.UserId(req.GetId()),
//		UserName:  req.GetUserName(),
//		FirstName: req.GetFirstName(),
//		LastName:  req.GetLastName(),
//	}
//	channel := t.serv.User().RetrieveOrCreate(&user, t.repo.User())
//
//	select {
//	case userRetrieved := <-channel:
//		if userRetrieved.Error != nil {
//			return nil, userRetrieved.Error
//		}
//		user := pb.User{
//			Id:        userRetrieved.User.Id.ToInt64(),
//			UserName:  userRetrieved.User.UserName,
//			FirstName: userRetrieved.User.FirstName,
//			LastName:  userRetrieved.User.LastName,
//		}
//		return &user, nil
//
//	case <-ctx.Done():
//		return nil, ErrorTimeOut
//	}
//}
//
//func (t tserver) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.Empty, error) {
//	user := domain.User{
//		Id:        domain.UserId(req.GetId()),
//		UserName:  req.GetUserName(),
//		FirstName: req.GetFirstName(),
//		LastName:  req.GetLastName(),
//	}
//	channel := t.serv.User().Update(&user, t.repo.User())
//
//	select {
//	case err := <-channel:
//		if err != nil {
//			return nil, err
//		}
//		return &pb.Empty{}, nil
//
//	case <-ctx.Done():
//		return nil, ErrorTimeOut
//	}
//}
//
//func (t tserver) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.Empty, error) {
//	userId := domain.UserId(req.GetId())
//	channel := t.serv.User().Delete(userId, t.repo.User())
//
//	select {
//	case err := <-channel:
//		if err != nil {
//			return nil, err
//		}
//		return &pb.Empty{}, nil
//
//	case <-ctx.Done():
//		return nil, ErrorTimeOut
//	}
//}

func (t tserver) RetrieveUser(ctx context.Context, req *pb.RetrieveUserRequest) (*pb.User, error) {
	userId := domain.UserId(req.GetId())
	userRetrieved := t.serv.User().Retrieve(userId, t.repo.User())
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
}

func (t tserver) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.Empty, error) {
	user := domain.User{
		Id:        domain.UserId(req.GetId()),
		UserName:  req.GetUserName(),
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
	}
	err := t.serv.User().Create(&user, t.repo.User())
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (t tserver) RetrieveOrCreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	user := domain.User{
		Id:        domain.UserId(req.GetId()),
		UserName:  req.GetUserName(),
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
	}
	userRetrieved := t.serv.User().RetrieveOrCreate(&user, t.repo.User())
	if userRetrieved.Error != nil {
		return nil, userRetrieved.Error
	}
	userR := pb.User{
		Id:        userRetrieved.User.Id.ToInt64(),
		UserName:  userRetrieved.User.UserName,
		FirstName: userRetrieved.User.FirstName,
		LastName:  userRetrieved.User.LastName,
	}
	return &userR, nil
}

func (t tserver) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.Empty, error) {
	user := domain.User{
		Id:        domain.UserId(req.GetId()),
		UserName:  req.GetUserName(),
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
	}
	err := t.serv.User().Update(&user, t.repo.User())
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (t tserver) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.Empty, error) {
	userId := domain.UserId(req.GetId())
	err := t.serv.User().Delete(userId, t.repo.User())
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
