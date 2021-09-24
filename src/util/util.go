package util

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/valyala/fasthttp"

	"github.com/golang-module/carbon"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
)

var Token = os.Getenv("VK_TOKEN")
var HypixelKey = os.Getenv("HYPIXEL_KEY")
var vk = api.NewVK(Token)
var NameRegex = regexp.MustCompile("^[a-zA-Z0-9_]{3,16}$")
var DB *sql.DB

var Client = fasthttp.Client{
	ReadTimeout: time.Second * 10,
}

func GetHypixelApi(method string, args string) (response string, err error) {
	code, res, err := Client.Get(nil, "https://api.hypixel.net/"+method+"?key="+HypixelKey+args)
	if err != nil {
		return
	}
	switch code {
		case 200:
			return string(res), err
		case 403:
			return "", errors.New("Invalid API key")
		case 429:
			return "", errors.New("Key throttle")
	}
	return
}

func GetUUID(name string) (mojang *Mojang, err error) {
	code, res, err := Client.Get(nil, "https://api.mojang.com/users/profiles/minecraft/"+name)
	if err != nil {
		return
	}
	if code != 200 {
		return nil, errors.New("Несуществующий ник")
	}

	mojang = &Mojang{}
	err = json.Unmarshal(res, &mojang)
	if err != nil {
		return
	}
	return
}

func GetName(uuid string) string {
	if uuid == "" {
		return "bebra"
	}
	_, res, err := Client.Get(nil, "https://api.mojang.com/user/profiles/"+uuid+"/names")
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

func MatchUsername(name string) bool {
	return NameRegex.Match([]byte(name))
}

func FormatTime(unix int) string {
	log.Println("unix", unix)
	lang := carbon.NewLanguage()
	diff := unix/1000 - int(time.Now().Unix())
	err := lang.SetLocale("ru")
	if err != nil {
		log.Fatal(err)
	}
	c := carbon.SetLanguage(lang)

	return c.Now().AddSeconds(diff).DiffForHumans()
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

func GetSlayerFromXp(xp int) (level int) {
	xp_table := Slayer_xp

	for i, x := range xp_table {
		if x > xp {
			break
		} else {
			level = i
		}
	}

	return level
}
