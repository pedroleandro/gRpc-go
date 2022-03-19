package services

import (
	"context"
	"fmt"
	"gRpc-go/pb"
	"io"
	"log"
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

func (*UserService) AddUsers(stream pb.UserService_AddUsersServer) error {
	users := []*pb.User{}

	for {
		request, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.Users{
				User: users,
			})
		}

		if err != nil {
			log.Fatalf("Erro ao receber dados de streaming: %v", err)
		}

		users = append(users, &pb.User{
			Id:    request.GetId(),
			Name:  request.GetName(),
			Email: request.GetEmail(),
		})

		fmt.Println("Adicionando usuario ", request.GetName())
	}
}

func (*UserService) AddUserStreamBoth(stream pb.UserService_AddUserStreamBothServer) error {
	for {
		request, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Erro ao receber dados de streaming do cliente: %v", err)
		}

		err = stream.Send(&pb.UserResultStream{
			Status: "Usuario adicionado",
			User:   request,
		})

		if err != nil {
			log.Fatalf("Erro ao enviar dados de streaming para o client: %v", err)
		}
	}
}
