package commands

import (
	"fmt"
	"hypixel-bot/cmd/util"
	"regexp"
)

var Bedwars = &util.Command{
	Name:      regexp.MustCompile("^(бв|бедварс|bw|bedwars)$"),
	Args:      1,
	ForAdmins: false,
	Trigger: func(name string, peer_id int, from_id int) (err error) {
		player, err := util.GetPlayer(name)
		if err != nil {
			return
		}
		message := fmt.Sprintf("Статистика игрока %s:\n\nУровень: %d\nW/L: %.2f (%.0f побед / %.0f поражений)\nFK/D: %.2f (%.0f убийств / %.0f смертей)\nBB/BL: %.2f (%.0f сломано / %.0f потеряно)\nWinstreak: %d",
			player.Displayname,
			player.Achievements.Bedwars_Level,
			player.Stats.Bedwars.Wins/player.Stats.Bedwars.Losses, player.Stats.Bedwars.Wins, player.Stats.Bedwars.Losses,
			player.Stats.Bedwars.Final_Kills/player.Stats.Bedwars.Final_Deaths, player.Stats.Bedwars.Final_Kills, player.Stats.Bedwars.Final_Deaths,
			player.Stats.Bedwars.Beds_Broken/player.Stats.Bedwars.Beds_Lost, player.Stats.Bedwars.Beds_Broken, player.Stats.Bedwars.Beds_Lost,
			player.Stats.Bedwars.Winstreak)

		err = util.SendMessage(peer_id, message)
		if err != nil {
			return
		}
		return
	},
}
