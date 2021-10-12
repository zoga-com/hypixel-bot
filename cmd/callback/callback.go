package callback

import (
	"context"
	"hypixel-bot/cmd/commands"
	"log"
	"net/http"

	"github.com/SevereCloud/vksdk/v2/callback"
	"github.com/SevereCloud/vksdk/v2/events"
)

func StartCallback() {
	log.Println("Started http server on port 80")

	cb := callback.NewCallback()

	cb.ConfirmationKey = ""
	cb.SecretKey = ""

	cb.MessageNew(func(ctx context.Context, obj events.MessageNewObject) {
		if obj.Message.PeerID != 2000000025 {
			log.Printf("%d: %s from %d", obj.Message.PeerID, obj.Message.Text, obj.Message.FromID)
			go commands.FindCommand(&obj)
		}
	})

	http.HandleFunc("/callback", cb.HandleFunc)

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}

}
