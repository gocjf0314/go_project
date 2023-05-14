package main

import (
	// "fmt"

	"log"
	"net"
	"os"

	"go_project/env"
	"go_project/handler"
	"go_project/proto/message"

	"google.golang.org/grpc"
)

type ClientServer struct {
	message.MessageServiceServer
}

func main() {
	env.LoadEnv()

	InitializeClientApp(os.Getenv("CLI_PORT"))
}

func InitializeClientApp(portNumber string) {
	// Initialize TCP connection
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		// TODO: Handle error...
		log.Fatalf("failed to listen: %v", err)
	}

	// Create new server
	grpcServer := grpc.NewServer()
	message.RegisterMessageServiceServer(grpcServer, &handler.MessageHandler{})

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		// TODO: Handle error...
		log.Fatalf("failed to serve: %s", err)
	}

	println("Run App....")
}
