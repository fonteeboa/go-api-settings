package main

import (
	"fmt"
	"golang-api-settings/internal/infra/http/routes"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Carregue as variáveis de ambiente a partir do arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	// Inicialize o roteador principal
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	// Passe o roteador para a função de configuração de rotas no pacote "routes"
	routes.ConfigureRoutes(r)

	err := r.Run(fmt.Sprintf("%s:%s", os.Getenv("GO_HOST"), os.Getenv("GO_PORT")))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server started at " + time.Now().String())
}
