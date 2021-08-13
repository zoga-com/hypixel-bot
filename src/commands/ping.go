package commands

import (
	"hypixel-bot/src/util"
)

var Debug = &util.Command{
	Name:      "ping",
	Args:      1,
	ForAdmins: false,
	Trigger: func(name string, peer_id int) (err error) {
		err = util.SendMessage(peer_id, "bebra")
		if err != nil {
			return
		}
		return
	},
}
