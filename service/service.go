package service

import (
	"context"
	"fmt"
	"go_project/database"
	"log"
	"strings"
)

type Server struct {
	ServiceInterfaceServer
}

func (s Server) GetData(ctx context.Context, req *GetMsg) (*GetResponse, error) {
	index := req.Index
	value, err := db.Get(fmt.Sprint(index))
	if err != nil {
		return nil, err
	}
	splited := strings.Split(value, "%")
	msg := MessageData{
		Index:   index,
		Content: splited[0],
	}
	return &GetResponse{MessageData: &msg}, nil
}

func (s Server) InsertData(ctx context.Context, res *InsertMsg) (*InsertReponse, error) {
	index, err := createMessage(res.MessageData)
	if err != nil || index == "" {
		return &InsertReponse{
			Index:      index,
			StatusCode: "500",
		}, err
	}

	return &InsertReponse{
		Index:      res.MessageData.Index,
		StatusCode: "200",
	}, nil
}

func createMessage(messageData *MessageData) (string, error) {
	index := messageData.Index
	content := messageData.Content

	_, err := db.Set(fmt.Sprint(index), content)
	if err != nil {
		log.Print(err.Error())
		return "", err
	}
	return index, nil
}

var db database.Database

func InitDB() {
	database, err := database.Factory("redis")
	if err != nil {
		log.Print(err)
		return
	}
	db = database
}

/* MySql */
// Create
// err := database.Create("`Messages` (`content`, `index`)", content, index)
// if err != nil {
// 	log.Fatalf(err.Error())
// 	return -1, err
// }
// response, err := Server.GetData(Server{}, context.Background(), &GetMsg{
// 	Index: messageData.Index,
// })
// if err != nil {
// 	log.Fatalf(err.Error())
// 	return -1, err
// }
// log.Printf("MessageData %v\n", response.GetMessageData())
// // log.Printf("Index: %d", msg.Index)
// // log.Printf("Content: %s", msg.Content)
// return response.GetMessageData().Index, nil
//
// Read
// columns := []string{"*"}
// table := "`Messages`"
// where := fmt.Sprintf("`index` = %d", index)
// rows := database.Read(columns, table, where)
// var msg MessageData
// // fmt.Print(rows.Columns())
// if rows.Next() {
// 	err := rows.Scan(&msg.Content, &msg.Index)
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 	}
// }
// return &GetResponse{MessageData: &msg}, nil
