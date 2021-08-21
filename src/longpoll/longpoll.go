package longpoll

import (
	"context"
	"database/sql"
	"hypixel-bot/src/commands"
	"hypixel-bot/src/util"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
)

func StartLongpoll() {
	vk := api.NewVK(util.Token)

	group, err := vk.GroupsGetByID(nil)
	if err != nil {
		log.Fatal(err)
	}

	lp, err := longpoll.NewLongPoll(vk, group[0].ID)
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("sqlite3", "./db/db.sql")
	if err != nil {log.Fatal(err)}

	lp.MessageNew(func(_ context.Context, obj events.MessageNewObject) {
		log.Printf("%d: %s from %d", obj.Message.PeerID, obj.Message.Text, obj.Message.FromID)
		commands.FindCommand(obj, db)
	})

	log.Println("Start Long Poll")
	if err := lp.Run(); err != nil {
		log.Fatal(err)
	}
}
