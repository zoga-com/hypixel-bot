package commands

import (
	"encoding/json"
	"fmt"
	"hypixel-bot/src/util"
	"strings"
	"sync"
	"time"
)

func buildAuctions(items *util.AuctionReturn, mojang *util.Mojang) string {
	text := make([]string, len(items.Auctions))

	text = append(text, "📗 Аукционы игрока "+mojang.Name+":\n")

	var wg sync.WaitGroup
	wg.Add(len(items.Auctions))

	for _, auc := range items.Auctions {
		go func(auc util.AuctionData) {
			defer wg.Done()
			if !auc.Claimed {
				fmt.Println(auc.GetHighestBid().Bidder)
				var icon, aucType string
				if int(time.Now().Unix()) > auc.End {
					icon = "✔"
				} else {
					icon = "⏳"
				}
				if auc.BIN {
					aucType = "BIN"
				} else {
					aucType = "Аукцион"
				}
				auction := fmt.Sprintf("%s [%s] %s (%s)\n• Истекает %s\n", icon, auc.Tier, auc.Name, aucType, util.FormatTime(auc.End))
				if auc.HighestBid == 0 {
					auction += fmt.Sprintf("• 💸 Начальная ставка: %d коинов\n", auc.StartingBid)
				} else {
					auction += fmt.Sprintf("• 💸 Последняя ставка: %d коинов\n• Ставка от игрока: %s\n", auc.HighestBid, util.GetName(auc.GetHighestBid().Bidder))
				}
				text = append(text, auction)
			}
		}(auc)
	}
	wg.Wait()
	return strings.Join(text, "\n\n")

}

var Auction = &util.Command{
	Name:      "ah",
	Args:      1,
	ForAdmins: false,
	Trigger: func(name string, peer_id int, from_id int) (err error) {
		mojang, err := util.GetUUID(name)
		if err != nil {
			return
		}
		res, err := util.GetHypixelApi("skyblock/auction", "&player="+mojang.Id)
		if err != nil {
			return
		}

		auctions := &util.AuctionReturn{}
		err = json.Unmarshal([]byte(res), &auctions)
		if err != nil {
			return
		}

		message := buildAuctions(auctions, mojang)

		err = util.SendMessage(peer_id, fmt.Sprintln(message))
		if err != nil {
			return
		}
		return
	},
}
