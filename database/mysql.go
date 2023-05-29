package database

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"os"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func LoadConnection() sql.DB {
// 	db, err := sql.Open(os.Getenv("DATABASE"), GetDSN())
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	return *db
// }

// func GetDSN() string {
// 	username := os.Getenv("DB_USER")
// 	password := os.Getenv("DB_PASSWORD")
// 	host := os.Getenv("DB_HOST")
// 	port := os.Getenv("DB_PORT")
// 	database := os.Getenv("DB_NAME")
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)

// 	return dsn
// }

// func Create(table string, values ...any) error {
// 	db := LoadConnection()
// 	defer db.Close()

// 	var temp string
// 	log.Print(temp)
// 	for i := 0; i < len(values); i++ {
// 		if i != 0 {
// 			temp = temp + ", ?"
// 		} else {
// 			temp = "?"
// 		}
// 	}

// 	query := fmt.Sprintf("INSERT INTO %s VALUES (%s)", table, temp)
// 	_, err := db.Query(query, values...)
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 		return err
// 	}
// 	return nil
// }

// func Update() {

// }

// func Read(columns []string, table string, where string) sql.Rows {
// 	db := LoadConnection()
// 	defer db.Close()

// 	var selection string
// 	for i := 0; i < len(columns); i++ {
// 		if i != 0 {
// 			selection = selection + ", " + columns[i]
// 		} else {
// 			selection = columns[i]
// 		}
// 	}

// 	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s", selection, table, where)
// 	result, err := db.Query(query)
// 	if err != nil {
// 		log.Fatalln(err.Error())
// 	}

// 	return *result
// }

// func Delete() {

// }
