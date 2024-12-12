package service

import (
	"context"
	v1 "userauth/api/user/v1"
	"userauth/internal/biz"
        "github.com/google/wire"
	"github.com/go-kratos/kratos/v2/log"
)

var ProviderSet = wire.NewSet(NewUserService)

type UserService struct {
	v1.UnimplementedUserServiceServer

	uc  *biz.UserUseCase
	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *UserService) AddUser(ctx context.Context, req *v1.AddUserRequest) (*v1.AddUserReply, error) {
	err := s.uc.AddUser(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return &v1.AddUserReply{Message: "User added successfully"}, nil
}

func (s *UserService) AuthenticateUser(ctx context.Context, req *v1.AuthenticateUserRequest) (*v1.AuthenticateUserReply, error) {
	authenticated, err := s.uc.AuthenticateUser(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return &v1.AuthenticateUserReply{Authenticated: authenticated}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserReply, error) {
	err := s.uc.DeleteUser(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteUserReply{Message: "User deleted successfully"}, nil
}
