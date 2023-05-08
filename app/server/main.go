package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

const portNumber = "9000"

/*
Network: tcp
Address: portNumber
*/

func main() {
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		// TODO: Handle error...
		log.Fatalf("failed to listen: %v", err)
	}

	// Create new Server
	grpcServer := grpc.NewServer()

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		// TODO: Handle error...
		log.Fatalf("failed to serve: %s", err)
	}
}
