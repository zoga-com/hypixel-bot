package callback

import (
	"context"
	"net/http"
	"hypixel-bot/src/commands"
	"log"

	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/callback"
)

func StartCallback() {
	log.Println("Started http server on port 80")

	cb := callback.NewCallback()

	cb.ConfirmationKey = ""
	cb.SecretKey = ""

	cb.MessageNew(func(ctx context.Context, obj events.MessageNewObject) {
		if obj.Message.PeerID != 2000000025 {
		log.Printf("%d: %s from %d", obj.Message.PeerID, obj.Message.Text, obj.Message.FromID)
		go commands.FindCommand(obj)
	}
	})

	http.HandleFunc("/callback", cb.HandleFunc)

	http.ListenAndServe(":80", nil)

}
