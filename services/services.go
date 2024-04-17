package services

import (
	"database/sql"
	"time"
)

var db *sql.DB
const Timeout = time.Second * 3

type Services struct {
	Item Item
	Manager Manager
	Department Department
}

func Register(dbPool *sql.DB) Services{
	db = dbPool
	return Services{}
}