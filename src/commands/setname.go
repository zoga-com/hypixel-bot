package commands

import (
	"database/sql"
	"strconv"
	"fmt"
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

		statement, err := db.Prepare(`INSERT INTO users (id, name, uuid) VALUES (?, ?, ?) ON CONFLICT(id) DO UPDATE SET id = ?, name = ?, uuid = ?`)
		if err != nil { return }

		_, err = statement.Exec(from_id, mojang.Name, mojang.Id, from_id, mojang.Name, mojang.Id)
		if err != nil { return }

        rows := db.QueryRow("SELECT * FROM users WHERE id =" + strconv.FormatInt(int64(from_id), 10))
        if err != nil {return}
		var id int
        var uuid string
        var username string

        err = rows.Scan(&id, &username, &uuid)
        if err != nil {return}
		util.SendMessage(peer_id, fmt.Sprintf("Вы теперь %s (UUID: %s)", username, uuid))

		return
	},
}
