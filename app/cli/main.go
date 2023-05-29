package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"go_project/env"
	"go_project/service"
	"go_project/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client service.ServiceInterfaceClient

func main() {
	env.LoadEnv()

	// server connector 생성
	serverConn := GetServerConnector()
	defer serverConn.Close()

	// client 생성
	client = service.NewServiceInterfaceClient(serverConn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for {
		// 메세지 입력
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Input message: ")

		input, _ := reader.ReadString('\n')
		splited := strings.Split(input, "")
		splited[len(splited)-1] = ""
		content := strings.Join(splited, "")

		// gRPC 통신에 쓸 GoChannel 생성
		server := make(chan string)

		index := utils.GetFormattedNow(time.Now().Local())
		log.Printf("Index: %s", index)

		go RequestServerToProcessMsg(ctx, server, &service.MessageData{
			Index:   index,
			Content: content,
		})

		// gRPC call 종료 후 channel을 통해 index 값 받기
		index = <-server

		// listener 양방향 통신
		ListenerConnHandler(index)
	}
}

func RequestServerToProcessMsg(ctx context.Context, server chan string, message *service.MessageData) {
	// 1. 서버로 메세지 저장 요청 후 응답 받기
	res, err := client.InsertData(ctx, &service.InsertMsg{
		MessageData: message,
	})
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	server <- res.Index
}

func ListenerConnHandler(index string) {
	// listener connector 객체 생성
	lisConn := GetListenerConnector()
	defer lisConn.Close()

	// socket 버퍼로 listner에게 index 전송
	lisConn.Write([]byte(index))

	// listener로 부터 버퍼로 응답 받기
	buffer := make([]byte, 1024)
	n, err := lisConn.Read(buffer)
	if err != nil {
		log.Fatalln("Occur Error during waiting response")
		log.Fatalf("Error[ListenerConnHandler]: %v", err)
	}

	// 응답 내용 할당 및 로깅
	message := string(buffer[:n])
	log.Printf("Recv from listener: %s", message)
}

func GetServerConnector() *grpc.ClientConn {
	host, portNumber := env.GetServerEnv()
	serverAddress := fmt.Sprintf("%s:%s", host, portNumber)
	log.Printf("Addr[server]: %s", host)
	log.Printf("Port[server]: %s", portNumber)

	var opts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial(serverAddress, opts...)
	if err != nil {
		log.Fatalln("Can not connect with server")
		log.Fatalf("Error[GetServerConnector]: %s", err.Error())
	}
	return conn
}

func GetListenerConnector() net.Conn {
	host, port := env.GetListenerEnv()
	listenerAddress := fmt.Sprintf("%s:%s", host, port)
	conn, err := net.Dial("tcp", listenerAddress)
	if err != nil {
		log.Fatalln("Can not connect with listener")
		log.Fatalf("Error[GetListenerConnector]: %s", err.Error())
	}
	return conn
}
