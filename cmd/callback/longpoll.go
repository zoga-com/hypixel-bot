package callback

import (
	"context"
	"hypixel-bot/cmd/commands"
	"hypixel-bot/cmd/util"
	"log"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
)

/*
 * Вы можете спросить: "Какого хуя лонгполл лежит в папке с названием каллбек?"
 * Отвечу на данный вопрос: я сам не ебу
 */

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
	lp.MessageNew(func(_ context.Context, obj events.MessageNewObject) {
		log.Printf("%d: %s from %d", obj.Message.PeerID, obj.Message.Text, obj.Message.FromID)
		commands.FindCommand(&obj)
	})

	log.Println("Start Long Poll")
	if err := lp.Run(); err != nil {
		log.Fatal(err)
	}
}
