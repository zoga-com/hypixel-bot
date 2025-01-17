package commands

import (
	"encoding/json"
	"fmt"
	"hypixel-bot/cmd/util"
	"regexp"
	"time"
)

var Where = &util.Command{
	Name:      regexp.MustCompile("where"),
	Args:      1,
	ForAdmins: false,
	Trigger: func(name string, peer_id int, from_id int) (err error) {
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
			message := fmt.Sprintf("Игрок %s был в сети %s (Сессия длилась %d / %d)", mojang.Name, time.Unix(status2.Player.LastLogout/1000, 0).Format("02.01.2006, в 15:04"), util.FormatTime2(int(status2.Player.LastLogout), int(status2.Player.LastLogin)), util.FormatTime(int(status2.Player.LastLogout)))
			err = util.SendMessage(peer_id, message)
		}
		return
	},
}
