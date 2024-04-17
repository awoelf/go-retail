package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq"
)

type Database struct {
	Client *sql.DB
}

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifetime = 5 * time.Minute

func Connect(dsn string) (*Database, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetConnMaxIdleTime(maxIdleDbConn)
	db.SetConnMaxLifetime(maxDbLifetime)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &Database{
		Client: db,
	}, nil
}