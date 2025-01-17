package commands

import (
	"errors"
	"fmt"
	"hypixel-bot/cmd/util"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/SevereCloud/vksdk/v2/events"
)

var commandRegex = regexp.MustCompile("^[./]([^ ]+)( .+)*$")
var Commands = []*util.Command{Bedwars, Skywars, Skyblock, Nick, Where}

func FindCommand(obj *events.MessageNewObject) {
	groups := commandRegex.FindStringSubmatch(obj.Message.Text)
	if groups == nil {
		return
	}
	var args []string = strings.Split(groups[2], " ")

	if groups != nil {
		if len(args) <= 2 {
			args = args[1:]
		} else {
			args = []string{}
		}

		command := struct {
			Name string
			Args int
		}{
			Name: groups[1],
			Args: len(args),
		}

		var wg sync.WaitGroup
		wg.Add(len(Commands))
		for _, it := range Commands {
			go func(it *util.Command) {
				var err error
				if !it.Name.Match([]byte(command.Name)) {
					return
				}

				if it.Args == command.Args {
					if util.MatchUsername(args[0]) {
						err = it.Trigger(args[0], obj.Message.PeerID, obj.Message.FromID)
					} else {
						err = errors.New("Несуществующий ник.")
					}
				} else if it.Args != 0 {
					row := util.DB.QueryRow("SELECT name FROM users WHERE id = ?", strconv.FormatInt(int64(obj.Message.FromID), 10))
					var name string
					err = row.Scan(&name)
					if name == "" || name == "false" || err != nil {
						_ = util.SendMessage(obj.Message.PeerID, "У вас не установлен ник.\nДля установки ника пропишите \"/name ник\"")
						return
					}
					err = it.Trigger(name, obj.Message.PeerID, obj.Message.FromID)
				}
				if err != nil {
					_ = util.SendMessage(obj.Message.PeerID, fmt.Sprintf("Произошла ошибка: %s (Игрока не существует?)", err))
				}
			}(it)
		}
	}
}
