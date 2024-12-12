package service

import (
	"context"

	pb "userauth/api/user/v1"
)

type UserServiceService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserServiceService() *UserServiceService {
	return &UserServiceService{}
}

func (s *UserServiceService) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserReply, error) {
	return &pb.AddUserReply{}, nil
}
func (s *UserServiceService) AuthenticateUser(ctx context.Context, req *pb.AuthenticateUserRequest) (*pb.AuthenticateUserReply, error) {
	return &pb.AuthenticateUserReply{}, nil
}
