package main

import (
	"context"
	"fmt"
	"gRpc-go/pb"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Nao foi possivel conectar ao servidor grpc: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	AddUserStreamBoth(client)
	//AddUsers(client)
	//AddUserVerbose(client)
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

func AddUsers(client pb.UserServiceClient) {
	request := []*pb.User{
		&pb.User{
			Id:    "1",
			Name:  "Pedro Leandro",
			Email: "pedro.leandrog@gmail.com",
		},

		&pb.User{
			Id:    "2",
			Name:  "Tiago Pereira",
			Email: "tiago.pereira@gmail.com",
		},

		&pb.User{
			Id:    "3",
			Name:  "Thiago da Cruz",
			Email: "thiago.cruz@gmail.com",
		},

		&pb.User{
			Id:    "4",
			Name:  "Gerson James",
			Email: "gerson.james@gmail.com",
		},

		&pb.User{
			Id:    "5",
			Name:  "Airton Sousa",
			Email: "airton.sousa@gmail.com",
		},
	}

	stream, err := client.AddUsers(context.Background())

	if err != nil {
		log.Fatalf("Erro ao enviar requisicao: %v", err)
	}

	for _, req := range request {
		fmt.Println("Enviando requisicao...")
		stream.Send(req)
		time.Sleep(time.Second * 2)
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Erro ao receber dados: %v", err)
	}

	fmt.Println(response)

}

func AddUserStreamBoth(client pb.UserServiceClient) {

	stream, err := client.AddUserStreamBoth(context.Background())

	if err != nil {
		log.Fatalf("Erro ao enviar requisicao: %v", err)
	}

	request := []*pb.User{
		&pb.User{
			Id:    "1",
			Name:  "Pedro Leandro",
			Email: "pedro.leandrog@gmail.com",
		},

		&pb.User{
			Id:    "2",
			Name:  "Tiago Pereira",
			Email: "tiago.pereira@gmail.com",
		},

		&pb.User{
			Id:    "3",
			Name:  "Thiago da Cruz",
			Email: "thiago.cruz@gmail.com",
		},

		&pb.User{
			Id:    "4",
			Name:  "Gerson James",
			Email: "gerson.james@gmail.com",
		},

		&pb.User{
			Id:    "5",
			Name:  "Airton Sousa",
			Email: "airton.sousa@gmail.com",
		},
	}

	wait := make(chan int)

	go func() {
		for _, req := range request {
			fmt.Println("Enviando usuario ", req.GetName())
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			response, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Erro ao receber dados: %v", err)
				break
			}

			fmt.Printf("Recebendo usuario %v com status %v\n", response.GetUser().GetName(), response.GetStatus())
		}
		close(wait)
	}()

	<-wait

}
