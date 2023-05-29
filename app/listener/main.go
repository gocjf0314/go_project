package main

import (
	"context"
	"fmt"
	"go_project/env"
	"go_project/service"
	"io"
	"log"
	"net"
)

func main() {
	env.LoadEnv()
	host, port := env.GetListenerEnv()
	listenerAddress := fmt.Sprintf("%s:%s", host, port)
	log.Printf("Connecting -> %s\n", listenerAddress)

	listener, err := net.Listen("tcp", listenerAddress)
	if err != nil {
		// TODO: Handle error...
		log.Fatalf("failed to listen: %v", err)
	}
	defer listener.Close()

	service.InitDB()

	println("Run Listener....")
	for {
		// 클라이언트 연결 수신
		fmt.Println("Waiting for connecting with client......")
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf(err.Error())
			return
		}
		defer conn.Close()

		go ConnecHandler(conn)
	}
}

func ConnecHandler(conn net.Conn) {
	// 양방향으로 cli와 통신할 수 있는 채널을 만든다.
	recvBuf := make([]byte, 4096) // receive buffer: 4kB
	for {
		n, err := conn.Read(recvBuf)
		if nil != err {
			if io.EOF == err {
				log.Printf("connection is closed from client; %v", conn.RemoteAddr().String())
				return
			}
			log.Printf("fail to receive data; err: %v", err)
			return
		}
		if 0 < n {
			data := recvBuf[:n]
			log.Println(string(data))
			index := string(data)

			message, err := service.Server{}.GetData(
				context.Background(),
				&service.GetMsg{Index: index},
			)
			if err != nil {
				log.Print(err.Error())
				return
			}
			log.Println(message)

			messageData := fmt.Sprintf("{index: %s, content: %s}", message.MessageData.Index, message.MessageData.Content)
			n, err := conn.Write([]byte(messageData))
			if err != nil {
				log.Fatalf("메시지 전송 중 오류가 발생했습니다: %v", err)
			}
			log.Printf("Data length: %d\n", n)
		}
	}
}
