package main

import (
	"database/sql"
	"hypixel-bot/cmd/callback"
	"hypixel-bot/cmd/util"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var err error
	util.DB, err = sql.Open("sqlite3", "./db/db.sql")
	if err != nil {
		log.Fatal(err)
	}
	callback.StartLongpoll()
}
