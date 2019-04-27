package handlers

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func init() {
	var err error
	db, err = sqlx.Connect("postgres", "postgres://postgres:postgres@localhost/golangw2?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}
