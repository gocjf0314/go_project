package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"go_project/env"
	"go_project/service"

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

		// 서버 통신(gRPC)
		server := make(chan int32)

		index := GetNowInt(time.Now().Local())
		log.Printf("Index: %d", index)

		message := service.MessageData{
			Index:   index,
			Content: content,
		}

		go RequestServerToProcessMsg(ctx, server, &message)

		index = <-server

		// listener 양방향 통신(TCP Socket)
		ListenerConnHandler(int(index))
	}
}

func RequestServerToProcessMsg(ctx context.Context, server chan int32, message *service.MessageData) {
	res, err := client.RequestServer(
		ctx,
		&service.RequestMsg{MessageData: message},
	)
	if err != nil {
		log.Print(err.Error())
		return
	}
	server <- res.Index
}

func ListenerConnHandler(index int) {
	// listener connector 객체 생성
	lisConn := GetListenerConnector()
	defer lisConn.Close()

	// socket 버퍼로 listner에게 index 전송
	value := strconv.Itoa(index)
	lisConn.Write([]byte(value))

	// listener로 부터 버퍼로 응답 받기
	buffer := make([]byte, 1024)
	n, err := lisConn.Read(buffer)
	if err != nil {
		log.Fatalf("메시지 수신 중 오류가 발생했습니다: %v", err)
	}

	// 응답 내용 할당 및 로깅
	message := string(buffer[:n])
	log.Printf("Recv from listener: %s", message)
}

func GetServerConnector() *grpc.ClientConn {
	host, portNumber := env.GetClientEnv()
	serverAddress := fmt.Sprintf("%s:%s", host, portNumber)
	log.Printf("server address: %s", host)
	log.Printf("server port number: %s", portNumber)

	var opts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial(serverAddress, opts...)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return conn
}

func GetListenerConnector() net.Conn {
	listenerAddress := fmt.Sprintf("%s:%s", os.Getenv("LOCALHOST"), os.Getenv("LISTENER_PORT"))
	conn, err := net.Dial("tcp", listenerAddress)
	if err != nil {
		log.Fatalf("서버에 연결할 수 없습니다: %v", err)
	}
	return conn
}

func GetNowInt(now time.Time) int32 {
	year := now.Year()
	month := ParseDigit(int(now.UTC().Month()))
	day := ParseDigit(now.Day())
	hour := ParseDigit(now.Hour())
	minute := ParseDigit(now.Minute())
	second := ParseDigit(now.Second())

	result := fmt.Sprintf("%d%s%s%s%s%s", year, month, day, hour, minute, second)
	index, err := strconv.Atoi(result)
	if err != nil {
		log.Fatal(err.Error())
	}
	return int32(index)
}

func ParseDigit(num int) string {
	if num < 10 {
		return fmt.Sprintf("0%d", num)
	}
	return fmt.Sprintf("%d", num)
}
