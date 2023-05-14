package handler

import (
	"context"
	"fmt"
	"go_project/database"
	"go_project/proto/message"
	"log"
)

type MessageHandler struct {
	message.MessageServiceServer
}

func (h MessageHandler) GetData(ctx context.Context, req *message.GetMsg) (*message.GetResponse, error) {
	index := req.Index

	log.Print(req)

	return &message.GetResponse{
		MessageData: &message.MessageData{
			Index:   index,
			Content: "Hello wellcome to gRPC world!",
		},
	}, nil
}

func (h MessageHandler) InsertData(ctx context.Context, res *message.InsertMsg) (*message.InsertReponse, error) {
	messageData := res.MessageData

	createMessage(messageData)

	return &message.InsertReponse{
		Index: messageData.Index,
	}, nil
}

func createMessage(messageData *message.MessageData) {
	index := messageData.Index
	content := fmt.Sprintf("%s(%d)", messageData.Content, index)

	database.Create("`Messages` (`index`, `content`)", index, content)
	// query := "INSERT INTO `Messages` (`index`, `content`) VALUES (?, ?)"
	// insertResult, err := db.ExecContext(context.Background(), query, index, content)
	// if err != nil {
	// 	log.Fatalf("impossible insert teacher: %s", err)
	// }

	// id, err := insertResult.LastInsertId()
	// if err != nil {
	// 	log.Fatalf("impossible to retrieve last inserted id: %s", err)
	// }
	// log.Printf("inserted id: %v", id)
}

func (h MessageHandler) RequestMsg(ctx context.Context, req *message.RequestData) (*message.ResponseData, error) {
	// 1. 서버로 메세지 저장 요청
	// 2. 요청에 대한 응답 받기
	res, err := h.InsertData(ctx, &message.InsertMsg{
		MessageData: req.MessageData,
	})

	// 3. status code와 index를 response에 할당
	if err != nil {
		log.Fatalf(err.Error())
		return &message.ResponseData{
			Index:      res.Index,
			StatusCode: "500",
		}, err
	}

	// Send index to listener
	defer h.SendData(ctx, req)

	return &message.ResponseData{
		Index:      res.Index,
		StatusCode: "200",
	}, nil
}

func (h MessageHandler) SendData(ctx context.Context, req *message.RequestData) (*message.ResponseData, error) {
	return &message.ResponseData{
		Index: req.MessageData.Index,
	}, nil
}

func (h MessageHandler) GetMsgFomrDB(ctx context.Context, req *message.ListenerReq) (*message.ListenerRes, error) {
	index := req.Index

	columns := []string{"*"}
	table := "`Messages`"
	where := fmt.Sprintf("index = %d", index)

	database.Read(columns, table, where)
	// db := database.LoadConnection()

	// query := "SELECT * FROM `Messages` WHERE index = ?"
	// result, err := db.QueryContext(context.Background(), query, index)
	// if err != nil {
	// 	log.Fatalf("impossible insert teacher: %s", err)
	// }

	// log.Printf("Result: %v", *result)

	return &message.ListenerRes{
		MessageData: &message.MessageData{
			Index:   index,
			Content: "djdjdjdjdjd",
		},
	}, nil
}
