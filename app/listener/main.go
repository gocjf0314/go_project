package main

import (
	"go_project/env"
	"go_project/handler"
	"go_project/proto/message"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	env.LoadEnv()

	InitializeLisenter(os.Getenv("LISTENER_PORT"))
}

/*
Network: tcp
Address: portNumber
*/
func InitializeLisenter(portNumber string) {
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
