package commands

import (
	"hypixel-bot/src/util"
)

var Ping = &util.Command{
	Name:      "ping",
	Args:      0,
	ForAdmins: false,
	Trigger: func(name string, peer_id int, from_id int) (err error) {
		err = util.SendMessage(peer_id, "bebra")
		if err != nil {
			return
		}
		return
	},
}
