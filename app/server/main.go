package main

import (
	"log"
	"net"
	"os"

	"go_project/env"
	"go_project/service"

	_ "github.com/go-sql-driver/mysql"

	"google.golang.org/grpc"
)

func InitializeServer(portNumber string) {
	// Initialize TCP connection
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		// TODO: Handle error...
		log.Fatalf("InitializeServer[net.Listen]]: %s", err.Error())
		return
	}

	// Create new server
	grpcServer := grpc.NewServer()
	service.RegisterServiceInterfaceServer(grpcServer, &service.Server{})
	log.Printf("Start gRPC server on %s port", portNumber)
	log.Printf("Running Server....")
	if err := grpcServer.Serve(lis); err != nil {
		// TODO: Handle error...
		log.Fatalf("InitializeServer[Serve]: %s", err)
	}
}

func main() {
	env.LoadEnv()

	InitializeServer(os.Getenv("SERVER_PORT"))
}
