package services

import (
	"context"
	"gRpc-go/pb"
)

//type UserServiceServer interface {
//	AddUser(context.Context, *User) (*User, error)
//	mustEmbedUnimplementedUserServiceServer()
//}

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(context context.Context, request *pb.User) (*pb.User, error) {
	return &pb.User{
		Id:    request.GetId(),
		Name:  request.GetName(),
		Email: request.GetEmail(),
	}, nil
}
