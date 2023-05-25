package service

import (
	"context"
	"fmt"
	"go_project/database"
	"log"
)

type Server struct {
	ServiceInterfaceServer
}

func (s Server) GetData(ctx context.Context, req *GetMsg) (*GetResponse, error) {
	index := req.Index
	log.Print(req)

	columns := []string{"*"}
	table := "`Messages`"
	where := fmt.Sprintf("`index` = %d", index)

	rows := database.Read(columns, table, where)
	var msg MessageData
	// fmt.Print(rows.Columns())
	if rows.Next() {
		err := rows.Scan(&msg.Content, &msg.Index)
		if err != nil {
			log.Fatalf(err.Error())
		}
	}

	return &GetResponse{MessageData: &msg}, nil
}

func (s Server) InsertData(ctx context.Context, res *InsertMsg) (*InsertReponse, error) {
	messageData := res.MessageData

	index, err := createMessage(messageData)
	if err != nil {
		return &InsertReponse{
			Index:      index,
			StatusCode: "500",
		}, err
	}

	return &InsertReponse{
		Index:      index,
		StatusCode: "200",
	}, nil
}

func (s Server) RequestServer(ctx context.Context, req *RequestMsg) (*IndexData, error) {
	// 1. 서버로 메세지 저장 요청
	// 2. 요청에 대한 응답 받기
	// 3. status code와 index를 response에 할당
	res, err := Server.InsertData(Server{}, ctx, &InsertMsg{
		MessageData: req.MessageData,
	})
	if err != nil {
		log.Fatalf(err.Error())
		return &IndexData{
			Index: -1,
		}, err
	}

	return &IndexData{Index: res.Index}, nil
}

func createMessage(messageData *MessageData) (int32, error) {
	index := messageData.Index
	content := messageData.Content

	err := database.Create("`Messages` (`content`, `index`)", content, index)
	if err != nil {
		log.Fatalf(err.Error())
		return -1, err
	}

	response, err := Server.GetData(Server{}, context.Background(), &GetMsg{
		Index: messageData.Index,
	})
	if err != nil {
		log.Fatalf(err.Error())
		return -1, err
	}

	log.Printf("MessageData %v\n", response.GetMessageData())
	// log.Printf("Index: %d", msg.Index)
	// log.Printf("Content: %s", msg.Content)
	return response.GetMessageData().Index, nil
}
