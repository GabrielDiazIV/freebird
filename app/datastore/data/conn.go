package data

import (
	"Freebird/app/system/env"
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	dbPool *pgxpool.Pool
	dbOnce sync.Once
)

func RwInstancePool() *pgxpool.Pool {
	dbOnce.Do(func() {
		connString := createConnectionString()

		db, err := pgxpool.Connect(context.Background(), connString)
		if err != nil {
			panic(err)
		}

		if err := db.Ping(context.Background()); err != nil {
			panic(err)
		}

		dbPool = db
	})

	return dbPool
}

func createConnectionString() string {
	c := env.GetConfig()
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.POSTGRES_USER,
		c.POSTGRES_PASS,
		c.POSTGRES_HOST,
		c.POSTGRES_PORT,
		"postgres",
	)
}
