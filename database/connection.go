package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func Connection() *pgx.Conn {
	instance, err := pgx.Connect(context.Background(), os.Getenv("MNSTR_DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	return instance
}
