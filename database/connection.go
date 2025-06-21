package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func Connection() (*pgx.Conn, error) {
	instance, err := pgx.Connect(context.Background(), os.Getenv("MNSTR_DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return instance, nil
}
