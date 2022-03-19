package main

import (
	"gRpc-go/pb"
	"gRpc-go/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Nao foi possivel conectar: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Nao foi possivel encontrar o servidor: %v", err)
	}

}
