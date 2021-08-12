package database

import (
	"log"

	"github.com/go-pg/pg/v10"
)

func New(databaseURL string) *pg.DB {
	opts, err := pg.ParseURL(databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	DB := pg.Connect(opts)

	return DB
}
