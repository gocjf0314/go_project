package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func LoadConnection() sql.DB {
	db, err := sql.Open(os.Getenv("DATABASE"), GetDSN())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	return *db
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

func Create(table string, values ...any) (int64, error) {
	db := LoadConnection()

	var temp string
	log.Print(temp)
	for i := 0; i < len(values); i++ {
		if i != 0 {
			temp = temp + ", ?"
		} else {
			temp = "?"
		}
	}

	query := fmt.Sprintf("INSERT INTO %s VALUES (%s)", table, temp)
	insertResult, err := db.ExecContext(context.Background(), query, values...)
	if err != nil {
		log.Fatalf("impossible insert teacher: %s", err)
		return -1, err
	}

	id, err := insertResult.LastInsertId()
	if err != nil {
		log.Fatalf("impossible to retrieve last inserted id: %s", err)
		return -1, err
	}

	log.Printf("inserted id: %v", id)
	return id, nil
}

func Update() {

}

func Read(columns []string, table string, where string) {
	db := LoadConnection()

	var selection string
	for i := 0; i < len(columns); i++ {
		if i != 0 {
			selection = selection + ", " + columns[i]
		} else {
			selection = columns[i]
		}
	}

	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s", selection, table, where)
	result, err := db.QueryContext(context.Background(), query)
	if err != nil {
		log.Fatalf("impossible insert teacher: %s", err)
	}

	log.Printf("Result: %v", *result)
}

func Delete() {

}
