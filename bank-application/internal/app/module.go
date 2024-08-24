package app

import (
	"database/sql"
	"log"
	"time"
)

type App struct {
	log *log.Logger
	db  *sql.DB
}

func NewApp(log *log.Logger, db *sql.DB) App {
	return App{
		log: log,
		db:  db,
	}
}

func (a *App) Close() {
	if err := a.db.Close(); err != nil {
		log.Printf("error while closing database connection: %s", err.Error())
	}

	// @Todo -> Closing log file is pending
}

type Account struct {
	Id         int64
	Owner      string
	Balance    int
	Currency   string
	CreateTime time.Time
}

type Entry struct {
	Id         int64
	AccountId  int
	Amount     int
	CreateTime time.Time
}

type Transfer struct {
	Id            int64
	FromAccountId int
	ToAccountId   int
	Amount        int
	CreateTime    time.Time
}
