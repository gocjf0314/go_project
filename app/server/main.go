package main

import (
	"log"
	"net"
	"os"

	"go_project/env"
	"go_project/handler"
	"go_project/proto/message"

	_ "github.com/go-sql-driver/mysql"

	"google.golang.org/grpc"
)

func main() {
	env.LoadEnv()

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
	message.RegisterMessageServiceServer(grpcServer, &handler.MessageHandler{})

	log.Printf("start gRPC server on %s port", portNumber)
	log.Printf("Running Server....")
	if err := grpcServer.Serve(lis); err != nil {
		// TODO: Handle error...
		log.Fatalf("failed to serve: %s", err)
	}
}

// func (s *MessageServer) GetData(ctx context.Context, req *message.GetMsg) (*message.GetResponse, error) {
// 	index := req.Index

// 	log.Print(req)

// 	return &message.GetResponse{
// 		MessageData: &message.MessageData{
// 			Index:   index,
// 			Content: "Hello wellcome to gRPC world!",
// 		},
// 	}, nil
// }

// func (s *MessageServer) InsertData(ctx context.Context, res *message.InsertMsg) (*message.InsertReponse, error) {
// 	messageData := res.MessageData

// 	createMessage(messageData)

// 	return &message.InsertReponse{
// 		Index: messageData.Index,
// 	}, nil
// }

// func createMessage(messageData *message.MessageData) {
// 	index := messageData.Index
// 	content := fmt.Sprintf("%s(%d)", messageData.Content, index)

// 	db := database.LoadConnection()

// 	query := "INSERT INTO `Messages` (`index`, `content`) VALUES (?, ?)"
// 	insertResult, err := db.ExecContext(context.Background(), query, index, content)
// 	if err != nil {
// 		log.Fatalf("impossible insert teacher: %s", err)
// 	}

// 	id, err := insertResult.LastInsertId()
// 	if err != nil {
// 		log.Fatalf("impossible to retrieve last inserted id: %s", err)
// 	}
// 	log.Printf("inserted id: %v", id)
// }
