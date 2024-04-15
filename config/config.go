package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect(dsn string) (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to database.")

	err = testDB(db)
	if err != nil {
		return nil, err
	}

	DB = db
	return DB, nil
}

func testDB(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Pinged database successfully.")

	return nil
}
