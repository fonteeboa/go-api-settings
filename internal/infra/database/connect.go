package database

import (
	"context"
	"fmt"
	"os"

	"github.com/vingarcia/ksql"
	"github.com/vingarcia/ksql/adapters/kpgx"
)

func NewConnect() (*ksql.DB, error) {
	ctx := context.Background()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		os.Getenv("DB_URI"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"),
	)

	db, err := kpgx.New(ctx, dsn, ksql.Config{})

	if err != nil {
		return nil, fmt.Errorf("unable to connect to database, %v", err)
	}

	fmt.Println("Connected to database")

	return &db, nil
}
