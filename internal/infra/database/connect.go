package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// NewConnect retorna uma instância de DBHandler, que é uma interface para manipular o banco de dados.
func NewConnect() (DBHandler, error) {
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

	// Conecte ao banco de dados usando o GORM
	gormDB, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	fmt.Println("Connected to the database")

	// Realiza a migração dos modelos de tabelas
	MigrateAllModels(&GormDBHandler{DB: gormDB})

	// Retorna uma instância de GormDBHandler que implementa a interface DBHandler
	return &GormDBHandler{DB: gormDB}, nil
}
