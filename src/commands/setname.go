package commands

import (
	"database/sql"
	"hypixel-bot/src/util"

	_ "github.com/mattn/go-sqlite3"
)

var Nick = &util.Command{
	Name:      "name",
	Args:      1,
	ForAdmins: false,
	Trigger: func(name string, peer_id int, from_id int) (err error) {
		mojang, err := util.GetUUID(name)
		if err != nil {
			return
		}

		db, err := sql.Open("sqlite3", "./db/db.sql")
		if err != nil {
			return
		}

		statement, err := db.Prepare(`INSERT INTO users (id, name) VALUES (?, ?) ON CONFLICT(id) DO UPDATE SET id = ?, name = ?`)
		if err != nil { return }

		_, err = statement.Exec(from_id, mojang.Name, from_id, mojang.Name)
		if err != nil { return }

		/*
        rows := db.QueryRow("SELECT name FROM users WHERE id =" + strconv.FormatInt(int64(from_id), 10))
        if err != nil {return}
        var username string

        err = rows.Scan(&username)
        if err != nil {return}
		util.SendMessage(peer_id, fmt.Sprintf("Вы теперь %s", username))
		*/

		return
	},
}
