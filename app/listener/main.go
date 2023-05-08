package listener

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

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

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		// TODO: Handle error...
		log.Fatalf("failed to serve: %s", err)
	}

	println("Run App....")
}

// func ListenClientRequest() (index int32)

// func SendIndexToServer(index int32) servicepb.GetResponse

func main() {
	err := godotenv.Load("go.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	InitializeLisenter(os.Getenv("LISTENER_PORT"))
}
