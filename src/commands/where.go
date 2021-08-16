package commands

import (
	"encoding/json"
	"fmt"
	"hypixel-bot/src/util"
	"time"
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
		res2, err := util.GetHypixelApi("player", "&uuid="+mojang.Id)
		if err != nil {
			return
		}
		status := &util.Status{}
		err = json.Unmarshal([]byte(res), &status)
		if err != nil {
			return
		}
		status2 := &util.Status2{}
		err = json.Unmarshal([]byte(res2), &status2)
		if err != nil {
			return
		}

		if status.Session.Online {
			message := fmt.Sprintf("Игрок %s онлайн на сервере %s, %s", mojang.Name, status.Session.GameType, status.Session.Mode)
			err = util.SendMessage(peer_id, message)
		} else {
			message := fmt.Sprintf("Игрок %s был в сети %s", mojang.Name, time.Unix(status2.Player.LastLogout / 1000, 0).Format("02.01.2006, 15:04"))
			err = util.SendMessage(peer_id, message)
		}
		return
	},
}
