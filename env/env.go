package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("env/go.env")
	if err != nil {
		log.Fatalf("Erro[LoadEnv]: %s", err)
	}
}

func GetServerEnv() (string, string) {
	host := os.Getenv("LOCALHOST")
	portNumber := os.Getenv("SERVER_PORT")
	return host, portNumber
}

func GetListenerEnv() (string, string) {
	host := os.Getenv("LOCALHOST")
	portNumber := os.Getenv("LISTENER_PORT")
	return host, portNumber
}
