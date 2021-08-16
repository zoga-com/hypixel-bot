package main

import (
	"context"
	"os"
        "bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"hypixel-bot/src/commands"
	"hypixel-bot/src/util"
	"log"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
)

func main() {
	vk := api.NewVK(util.Token)
	q, _ := os.Open("../resources/quote.png")
	quote, _, _ := image.Decode(q)

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

		if obj.Message.FromID == 351445954 && obj.Message.Text != "" {
			fontBytes, _ := ioutil.ReadFile("../resources/Ubuntu.ttf")
			f, _ := truetype.Parse(fontBytes)
			opts := truetype.Options{}
			opts.Size = 35
			fg := color.RGBA{255, 255, 255, 255}
			rgba := image.NewRGBA(image.Rect(0, 0, 1000, 500))
			draw.Draw(rgba, quote.Bounds(), quote, image.Point{}, draw.Over)
			c := freetype.NewContext()
			c.SetFont(f)
			c.SetFontSize(opts.Size)
			c.SetClip(rgba.Bounds())
			c.SetDst(rgba)
			c.SetSrc(&image.Uniform{fg})

			pt := freetype.Pt(100, 190)// draw the string
			_, err = c.DrawString("«" + obj.Message.Text + "»", pt)

        		buffer := &bytes.Buffer{}
        		err = png.Encode(buffer, rgba)
        		if err != nil {return}
		
        		photosPhoto, err := vk.UploadMessagesPhoto(198657266, buffer)
        		if err != nil {
            		log.Fatalf("Error uploading photo: %s", err)
        		}
		
        		builder := params.NewMessagesSendBuilder()
        		builder.RandomID(0)
			//builder.ContentSource(fmt.Sprintf(`{"owner_id":%d, "peer_id":%d, "conversation_message_id":%d}`, obj.Message.FromID, obj.Message.PeerID, obj.Message.ConversationMessageID))
			builder.PeerIDs([]int{obj.Message.PeerID})
        		builder.Attachment(fmt.Sprintf("photo%d_%d_%s", photosPhoto[0].OwnerID, photosPhoto[0].ID, photosPhoto[0].AccessKey))
        		_, err = vk.MessagesSendUserIDs(builder.Params)
        		if err != nil {
            		log.Fatal(err)
        		}

		}

		commands.FindCommand(obj)
	})

	log.Println("Start Long Poll")
	if err := lp.Run(); err != nil {
		log.Fatal(err)
	}
}
