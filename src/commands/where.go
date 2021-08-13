package commands

import (
	"encoding/json"
	"fmt"
	"hypixel-bot/src/util"
)

var Where = &util.Command{
	Name:      "where",
	Args:      1,
	ForAdmins: false,
	Trigger: func(name string, peer_id int) (err error) {
		mojang, err := util.GetUUID(name)
		if err != nil {
			return
		}
		res, err := util.GetHypixelApi("status", "&uuid="+mojang.Id)
		if err != nil {
			return
		}
		status := &util.Status{}
		err = json.Unmarshal([]byte(res), &status)
		if err != nil {
			return
		}

		if status.Session.Online {
			message := fmt.Sprintf("Игрок %s онлайн на сервере %s, %s", mojang.Name, status.Session.GameType, status.Session.Mode)
			err = util.SendMessage(peer_id, message)
		} else {
			err = util.SendMessage(peer_id, "Игрок не онлайн.")
		}
		return
	},
}
