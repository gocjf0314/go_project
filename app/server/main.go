package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	"go_project/env"
	"go_project/proto/message"

	_ "github.com/go-sql-driver/mysql"

	"google.golang.org/grpc"
)

type ServiceServer struct {
	message.ServiceInterfaceServer
}

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
	message.RegisterServiceInterfaceServer(grpcServer, &ServiceServer{})

	log.Printf("start gRPC server on %s port", portNumber)
	log.Printf("Running Server....")
	if err := grpcServer.Serve(lis); err != nil {
		// TODO: Handle error...
		log.Fatalf("failed to serve: %s", err)
	}
}

func GetDSN() string {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)

	return dsn
}

func LoadDatabase() sql.DB {
	db, err := sql.Open(os.Getenv("DATABASE"), GetDSN())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	return *db
}

func (s *ServiceServer) GetData(ctx context.Context, req *message.GetMsg) (*message.GetResponse, error) {
	index := req.Index

	log.Print(req)

	return &message.GetResponse{
		MessageData: &message.MessageData{
			Index:   index,
			Content: "Hello wellcome to gRPC world!",
		},
	}, nil
}

func (s *ServiceServer) InsertData(ctx context.Context, res *message.InsertMsg) (*message.InsertReponse, error) {
	messageData := res.MessageData
	index := messageData.Index
	content := fmt.Sprintf("%s(%d)", messageData.Content, index)

	db := LoadDatabase()

	query := "INSERT INTO `Messages` (`index`, `content`) VALUES (?, ?)"
	insertResult, err := db.ExecContext(context.Background(), query, index, content)
	if err != nil {
		log.Fatalf("impossible insert teacher: %s", err)
	}

	id, err := insertResult.LastInsertId()
	if err != nil {
		log.Fatalf("impossible to retrieve last inserted id: %s", err)
	}
	log.Printf("inserted id: %v", id)

	return &message.InsertReponse{
		Index: index,
	}, nil
}

// func listenClientRequest()
// func (s *ServiceServer) InsertMsg(msg service.InsertMsg) (*service.InsertReponse, error)
// func createMessage(msg service.MessageData)
