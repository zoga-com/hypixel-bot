package commands

import (
	"fmt"
	"hypixel-bot/src/util"
	"regexp"
	"sync"

	"github.com/SevereCloud/vksdk/v2/events"
)

var commandRegex = regexp.MustCompile("^[./]([^ ]+)( .+)*$")
var Commands = []*util.Command{Bedwars, Skywars, Where, Auction, Skyblock}

func FindCommand(obj events.MessageNewObject) {
	groups := commandRegex.FindStringSubmatch(obj.Message.Text)
	var args []string

	if groups != nil {
		if len(regexp.MustCompile(" ").Split(groups[2], -1)) <= 2 {
			args = regexp.MustCompile(" ").Split(groups[2], -1) // самый отвратительный код в проекте
		} else {
			args = []string{}
		}

		command := &util.Command{
			Name: groups[1],
			Args: len(args) - 1,
		}

		var wg sync.WaitGroup
		wg.Add(len(Commands))
		for _, it := range Commands {
			go func(it *util.Command) {
				if it.Name != command.Name {
					return
				}

				if it.Args == command.Args {
					err := it.Trigger(args[1], obj.Message.PeerID, obj.Message.FromID)
					if err != nil {
						util.SendMessage(obj.Message.PeerID, fmt.Sprintf("Произошла ошибка: %s (Игрока не существует?)", err))
					}
				}
			}(it)
		}
	}
}
