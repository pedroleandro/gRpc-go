package main

import (
	"context"
	"fmt"
	"gRpc-go/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Nao foi possivel conectar ao servidor grpc: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	AddUser(client)
}

func AddUser(client pb.UserServiceClient) {
	request := &pb.User{
		Id:    "1",
		Name:  "Pedro Leandro",
		Email: "pedro.leandrog@gmail.com",
	}

	response, err := client.AddUser(context.Background(), request)

	if err != nil {
		log.Fatalf("Nao foi possivel realizar a requisicao grpc: %v", err)
	}

	fmt.Println(response)
}
