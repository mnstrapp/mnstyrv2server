package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var instance *pgx.Conn

func Connection() *pgx.Conn {
	if instance != nil {
		return instance
	}

	var err error
	instance, err = pgx.Connect(context.Background(), os.Getenv("MNSTR_DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	return instance
}
