package utils

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewDbConnection(log *log.Logger) (db *sql.DB) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/bank")
	if err != nil {
		log.Panic(err.Error())
		return
	}
	log.Print("database connection successful")
	return
}
