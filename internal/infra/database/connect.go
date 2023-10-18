package database

import (
	"context"
	"fmt"
	"os"
	"github.com/vingarcia/ksql"
	"github.com/vingarcia/ksql/adapters/kpgx"
)

func NewConnect() (ksql.DB, error) {
	// Leitura das variáveis de ambiente do arquivo .env
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_EXTERNAL_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

    if host == "" {
        host = "127.0.0.1"
    }

	// Construa a string de conexão
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Crie um contexto
	ctx := context.Background()

	// Conecte ao banco de dados usando o adaptador kpgx
	db, err := kpgx.New(ctx, dsn, ksql.Config{})
	if err != nil {
		return ksql.DB{}, fmt.Errorf("unable to connect to database: %v", err)
	}
    fmt.Println("Connected to the database")

    return db, nil
}
