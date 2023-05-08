package server

import (
	"log"
	"net"
	"os"

	servicepb "go_project/service"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type myServer struct {
	servicepb.ServiceInterfaceServer
}

func main() {
	err := godotenv.Load("go.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	InitializeGRPCServer(os.Getenv("SERVER_PORT"))
}

func InitializeGRPCServer(portNumber string) {
	// Initialize TCP connection
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		// TODO: Handle error...
		log.Fatalf("failed to listen: %v", err)
	}

	// Create new server
	grpcServer := grpc.NewServer()

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		// TODO: Handle error...
		log.Fatalf("failed to serve: %s", err)
	}

	println("Run App....")
}

// func listenClientRequest {}

// func (s *myServer) InsertMsg(msg servicepb.InsertMsg) (*servicepb.InsertReponse, error) {

// }

func (s *myServer) GetData(msg servicepb.GetMsg) (*servicepb.GetResponse, error) {
	var index int32 = msg.Index

	var messageData *servicepb.GetResponse

	*messageData = GetMessageFromDB(index)

	return messageData, nil
}

func GetMessageFromDB(index int32) servicepb.GetResponse {
	// var msg servicepb.GetResponse = ......

	var msgData = servicepb.MessageData{
		Index:   12,
		Content: "Hello, Wellcome to gRPC world!",
	}

	return servicepb.GetResponse{
		MessageData: &msgData,
	}
}

// func sendDataTo(index int32, receiver) {}
