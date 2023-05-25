package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("env/go.env")
	if err != nil {
		log.Fatal("Error loading .env file")
		log.Fatal(err)
	}
}

func GetClientEnv() (string, string) {
	host := os.Getenv("LOCALHOST")
	portNumber := os.Getenv("SERVER_PORT")
	return host, portNumber
}
