package main

import (
    "log"
    "os"
    "fmt"
	"time"
    "github.com/joho/godotenv"
    "github.com/gin-gonic/gin"
    "golang-api-settings/internal/infra/http/routes"
)

func main() {

	// Carregue as variáveis de ambiente a partir do arquivo .env
    if err := godotenv.Load(); err != nil {
        log.Fatal("Erro ao carregar o arquivo .env")
    }

    // Inicialize o roteador principal
    r := gin.Default()

    // Passe o roteador para a função de configuração de rotas no pacote "routes"
    routes.ConfigureRoutes(r)

    err := r.Run(fmt.Sprintf("%s:%s", os.Getenv("GO_HOST"), os.Getenv("GO_PORT")));

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Server started at " + time.Now().String())
}
