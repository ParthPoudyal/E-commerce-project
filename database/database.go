package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Db *pgxpool.Pool
}

var Db Database

func Connectdb() error {

	dsn := os.Getenv("DB_DSN")

	if dsn == "" {
		return &DBerror{message: "DB_DSN not set"}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return &DBerror{message: "Error Connecting to DB", cause: err}
	}

	if err = pool.Ping(ctx); err != nil {
		return &DBerror{message: "Unable to Ping DB", cause: err}
	}

	fmt.Println("DATABASE CONNECTED")

	Db = Database{
		Db: pool,
	}

	return nil
}
