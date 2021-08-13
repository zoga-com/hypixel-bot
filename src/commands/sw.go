package commands

import (
	"fmt"
	"hypixel-bot/src/util"
)

var Skywars = &util.Command{
	Name:      "sw",
	Args:      1,
	ForAdmins: false,
	Trigger: func(name string, peer_id int, from_id int) (err error) {
		player, err := util.GetPlayer(name)
		if err != nil {
			return
		}
		message := fmt.Sprintf("Статистика игрока %s\n\nУровень: %s\nW/L: %.2f (%.0f побед / %.0f поражений)\nK/D: %.2f (%.0f убийств / %.0f смертей)\nМонет: %d",
			player.Displayname,
			player.Stats.SkyWars.Level[3:],
			player.Stats.SkyWars.Wins/player.Stats.SkyWars.Losses, player.Stats.SkyWars.Wins, player.Stats.SkyWars.Losses,
			player.Stats.SkyWars.Kills/player.Stats.SkyWars.Deaths, player.Stats.SkyWars.Kills, player.Stats.SkyWars.Deaths,
			player.Stats.SkyWars.Coins)

		err = util.SendMessage(peer_id, message)
		return
	},
}
