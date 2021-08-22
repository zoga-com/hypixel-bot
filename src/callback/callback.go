package longpoll

import (
	"context"
	"database/sql"
	"net/http"
	"hypixel-bot/src/commands"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/callback"
)

func StartLongpoll() {
	log.Println("Started http server on port 80")
	db, err := sql.Open("sqlite3", "./db/db.sql")
	if err != nil {log.Fatal(err)}

	cb := callback.NewCallback()

	cb.ConfirmationKey = ""
	cb.SecretKey = ""

	cb.MessageNew(func(ctx context.Context, obj events.MessageNewObject) {
		if obj.Message.PeerID != 2000000025 {
		log.Printf("%d: %s from %d", obj.Message.PeerID, obj.Message.Text, obj.Message.FromID)
		go commands.FindCommand(obj, db)
	}
	})

	http.HandleFunc("/callback", cb.HandleFunc)

	http.ListenAndServe(":80", nil)

}
