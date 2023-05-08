package cli

import (
	// "fmt"
	"log"
	"net"
	"os"

	service "go_project/service"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

/*
Network: tcp
Address: portNumber
*/
func InitializeClientApp(portNumber string) {
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

func requestMsg(msg service.InsertMsg) {

}

func sendIndexToLisener(index int32) {

}

func main() {
	err := godotenv.Load("go.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	InitializeClientApp(os.Getenv("CLI_PORT"))

	// Create Golang Channel
	ch := make(chan int32)

	// Channel get data from server
	// go requestMsg()

	var response int32 = <-ch

	sendIndexToLisener(response)
}
