package commands

import (
	"errors"
	"fmt"
	"hypixel-bot/src/util"
	"log"
	"regexp"
	"strconv"
	"sync"

	"github.com/SevereCloud/vksdk/v2/events"
)

var commandRegex = regexp.MustCompile("^[./]([^ ]+)( .+)*$")
var Commands = []*util.Command{Bedwars, Skywars, Skyblock, Nick, Ping, Auction}

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
				if !regexp.MustCompile(it.Name).Match([]byte(command.Name)) {
					return
				}

				if it.Args == command.Args {
					var err error
					if util.MatchUsername(args[1]) {
						err = it.Trigger(args[1], obj.Message.PeerID, obj.Message.FromID)
					} else { err = errors.New("Несуществующий ник.") }
					if err != nil {
						util.SendMessage(obj.Message.PeerID, fmt.Sprintf("Произошла ошибка: %s (Игрока не существует?)", err))
					}
				} else {
					row := util.DB.QueryRow("SELECT name FROM users WHERE id=" + strconv.FormatInt(int64(obj.Message.FromID), 10))
					var name string
					err := row.Scan(&name)
					if name == "" || name == "false" { util.SendMessage(obj.Message.PeerID, "У вас не установлен ник.\nДля установки ника пропишите \"/name ник\""); return }
					if err != nil { log.Fatal(err) }
					it.Trigger(name, obj.Message.PeerID, obj.Message.FromID)
				}
			}(it)
		}
	}
}
