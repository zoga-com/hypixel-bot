package util

import (
	"encoding/json"
	"log"
	"os"

	"github.com/valyala/fasthttp"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
)

var Token = os.Getenv("VK_TOKEN")
var HypixelKey = os.Getenv("HYPIXEL_KEY")
var vk = api.NewVK(Token)

func GetHypixelApi(method string, args string) (response string, err error) {
	_, res, err := fasthttp.Get(nil, "https://api.hypixel.net/"+method+"?key="+HypixelKey+args)
	if err != nil {
		return
	}
	return string(res), err
}

func GetUUID(name string) (mojang *Mojang, err error) {
	_, res, err := fasthttp.Get(nil, "https://api.mojang.com/users/profiles/minecraft/"+name)

	if err != nil {
		return
	}

	mojang = &Mojang{}
	err = json.Unmarshal(res, &mojang)
	if err != nil {
		return
	}
	return
}

func GetName(uuid string) string {
	_, res, err := fasthttp.Get(nil, "https://api.mojang.com/user/profiles/"+uuid+"/names")
	if err != nil {
		log.Fatal(err)
	}

	name := []Name{}
	_ = json.Unmarshal([]byte(res), &name)

	return name[len(name)-1].Name
}

func GetPlayer(name string) (response Player, err error) {
	uuid, err := GetUUID(name)
	if err != nil {
		return
	}
	res, err := GetHypixelApi("player", "&uuid="+uuid.Id)
	if err != nil {
		return
	}
	player := &PlayerResponse{}
	err = json.Unmarshal([]byte(res), &player)
	if err != nil {
		return
	}
	return player.Player, err
}

func SendMessage(peer_id int, message string) (err error) {
	msg := params.NewMessagesSendBuilder()
	msg.Message(message)
	msg.RandomID(0)
	msg.PeerID(peer_id)

	_, err = vk.MessagesSend(msg.Params)
	if err != nil {
		return
	}
	return
}
