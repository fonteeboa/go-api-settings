package providers

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	var err error
	currentWorkDirectory, _ := os.Getwd()
	err = godotenv.Load(fmt.Sprintf("%s/.env", currentWorkDirectory))
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
}
