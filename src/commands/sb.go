package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"io/ioutil"
	"regexp"
	"strings"
	"sync"

	"hypixel-bot/src/util"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

var Skyblock = &util.Command{
	Name:      regexp.MustCompile("^(skyblock|сб|скайблок|sb)$"),
	Args:      1,
	ForAdmins: false,
	Trigger: func(name string, peer_id int, from_id int) (err error) {
		mojang, err := util.GetUUID(name)
		if err != nil {
			return
		}
		profile := &util.Profile{}

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			_, res, err := util.Client.Get(nil, "https://api.slothpixel.me/api/skyblock/profile/"+mojang.Id+"?key="+util.HypixelKey)
			if err != nil {
				return
			}
			err = json.Unmarshal([]byte(res), &profile)
			if err != nil {
				return
			}
			wg.Done()
		}()
		var skin image.Image
		go func() {
			_, res, err := util.Client.Get(nil, "https://visage.surgeplay.com/full/448/"+mojang.Id)
			if err != nil {
				return
			}
			skin, _, err = image.Decode(bytes.NewReader(res))
			if err != nil {
				return
			}
			wg.Done()
		}()

		wg.Wait()

		if err != nil {
			return
		}
		member := profile.Members[mojang.Id]

		message := fmt.Sprintf("Статистика игрока %s на Skyblock\nПрофиль: %s\n\nHP: %.2f | EHP: %.2f\nDefense: %.2f | True Defense: %.2f\nStrength: %.2f\nAttack Speed: %.0f%%\nCrit Chance: %.2f%% | Crit Damage: %.2f%%\nIntelligence: %.2f\n\nPurse: %.2f\nBank: %.2f\n\nAverage: %.2f\n\nTaming level: %d/50 (%.0f%%)\nFarming level: %d/60 (%.0f%%)\nMining level: %d/60 (%.0f%%)\nCombat level: %d/60 (%.0f%%)\nForaging level: %d/50 (%.0f%%)\nFishing level: %d/50 (%.0f%%)\nEnchanting level: %d/60 (%.0f%%)\nAlchemy level: %d/50 (%.0f%%)\nCarpentry level: %d/50 (%.0f%%)\nRunecrafting level: %d/25 (%.0f%%)\n\nZombie | Spider | Wolf | Enderman\nВсего XP: %d\nУровень: %d | %d | %d | %d\nXP: %d | %d | %d | %d\n ",
			mojang.Name,
			profile.CuteName,
			member.Attributes.Health,
			member.Attributes.EffectiveHealth,
			member.Attributes.Defense,
			member.Attributes.TrueDefense,
			member.Attributes.Strength,
			member.Attributes.BonusAttackSpeed,
			member.Attributes.CritChance,
			member.Attributes.CritDamage,
			member.Attributes.Intelligence,
			member.Purse,
			profile.Banking.Balance,
			member.Average,
			member.Skills.Taming.Level, member.Skills.Taming.Progress*100,
			member.Skills.Farming.Level, member.Skills.Farming.Progress*100,
			member.Skills.Mining.Level, member.Skills.Mining.Progress*100,
			member.Skills.Combat.Level, member.Skills.Combat.Progress*100,
			member.Skills.Foraging.Level, member.Skills.Foraging.Progress*100,
			member.Skills.Fishing.Level, member.Skills.Fishing.Progress*100,
			member.Skills.Enchanting.Level, member.Skills.Enchanting.Progress*100,
			member.Skills.Alchemy.Level, member.Skills.Alchemy.Progress*100,
			member.Skills.Carpentry.Level, member.Skills.Carpentry.Progress*100,
			member.Skills.Runecrafting.Level, member.Skills.Runecrafting.Progress*100,
			member.Slayers.Zombie.Xp+member.Slayers.Enderman.Xp+member.Slayers.Spider.Xp+member.Slayers.Wolf.Xp,
			util.GetSlayerFromXp(member.Slayers.Zombie.Xp), util.GetSlayerFromXp(member.Slayers.Spider.Xp), util.GetSlayerFromXp(member.Slayers.Wolf.Xp), util.GetSlayerFromXp(member.Slayers.Enderman.Xp),
			member.Slayers.Zombie.Xp, member.Slayers.Spider.Xp, member.Slayers.Wolf.Xp, member.Slayers.Enderman.Xp)
		fontBytes, _ := ioutil.ReadFile("../resources/Ubuntu.ttf")
		f, _ := truetype.Parse(fontBytes)
		opts := truetype.Options{}
		opts.Size = 12
		bg, fg := color.RGBA{46, 52, 64, 255}, color.RGBA{129, 161, 193, 255}
		rgba := image.NewRGBA(image.Rect(0, 0, 600, 600))
		draw.Draw(rgba, rgba.Bounds(), &image.Uniform{bg}, image.Point{}, draw.Src)
		c := freetype.NewContext()
		c.SetFont(f)
		c.SetFontSize(opts.Size)
		c.SetClip(rgba.Bounds())
		c.SetDst(rgba)
		c.SetSrc(&image.Uniform{fg})
		text := strings.Split(message, "\n")
		pt := freetype.Pt(300, 10+int(c.PointToFixed(20)>>6))
		for _, s := range text {
			_, err = c.DrawString(s, pt)
			if err != nil {
				return
			}
			pt.Y += c.PointToFixed(opts.Size * 1.5)
		}

		sp2 := image.Point{skin.Bounds().Dx() - 250, 75}
		r2 := image.Rectangle{sp2, sp2.Add(skin.Bounds().Size())}
		draw.Draw(rgba, r2, skin, image.Point{0, 0}, draw.Over)

		buffer := &bytes.Buffer{}
		err = util.Encoder.Encode(buffer, rgba)
		if err != nil {
			return
		}

		photosPhoto, err := util.VK.UploadMessagesPhoto(198657266, buffer)
		if err != nil {
			return
		}

		builder := params.NewMessagesSendBuilder()
		builder.RandomID(0)
		builder.PeerID(peer_id)
		builder.Attachment(fmt.Sprintf("photo%d_%d_%s", photosPhoto[0].OwnerID, photosPhoto[0].ID, photosPhoto[0].AccessKey))
		_, err = util.VK.MessagesSend(builder.Params)
		if err != nil {
			return
		}
		return
	},
}
