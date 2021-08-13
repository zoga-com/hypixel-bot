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

	text = append(text, "📗 Аукционы игрока "+mojang.Name+":")

	var wg sync.WaitGroup
	wg.Add(len(items.Auctions))

	for _, auc := range items.Auctions {
		go func(auc util.AuctionData) {
			defer wg.Done()
			if !auc.Claimed {
				var icon, aucType string
				if int(time.Now().Unix()) < auc.End {
					icon = "✔"
				} else {
					icon = "🚫"
				}
				if !auc.BIN {
					aucType = "BIN"
				} else {
					aucType = "Аукцион"
				}
				text = append(text, fmt.Sprintf("%s [%s] Shredder\n• 💸 Последняя ставка: %d коинов\n• Ставка от игрока: %s\n• 💭 Тип: %s",
					icon,
					auc.Tier,
					auc.HighestBid,
					util.GetName(auc.GetHighestBid().Bidder),
					aucType))
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
	Trigger: func(name string, peer_id int) (err error) {
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
