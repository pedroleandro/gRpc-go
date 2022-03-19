package main

import (
	"context"
	"fmt"
	"gRpc-go/pb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Nao foi possivel conectar ao servidor grpc: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	AddUserVerbose(client)
	//AddUser(client)
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

func AddUserVerbose(client pb.UserServiceClient) {
	request := &pb.User{
		Id:    "1",
		Name:  "Pedro Leandro",
		Email: "pedro.leandrog@gmail.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), request)

	if err != nil {
		log.Fatalf("Nao foi possivel realizar a requisicao grpc: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Nao foi possivel receber os dados a requisicao grpc: %v", err)
		}

		fmt.Println("Status: ", stream.Status)
	}
}
