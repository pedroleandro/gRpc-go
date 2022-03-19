package services

import (
	"context"
	"gRpc-go/pb"
	"time"
)

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

func (*UserService) AddUserVerbose(request *pb.User, stream pb.UserService_AddUserVerboseServer) error {
	stream.Send(&pb.UserResultStream{
		Status: "iniciando",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 2)

	stream.Send(&pb.UserResultStream{
		Status: "Inserindo",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Cadastrado com sucesso",
		User: &pb.User{
			Id:    request.GetId(),
			Name:  request.GetName(),
			Email: request.GetEmail(),
		},
	})

	time.Sleep(time.Second * 5)

	stream.Send(&pb.UserResultStream{
		Status: "Finalizado",
		User: &pb.User{
			Id:    request.GetId(),
			Name:  request.GetName(),
			Email: request.GetEmail(),
		},
	})

	time.Sleep(time.Second * 2)

	return nil
}
