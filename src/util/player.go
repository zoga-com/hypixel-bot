package util

import "time"

// huge fucking player struct

type PlayerResponse struct {
	Success bool   `json:"success"`
	Player  Player `json:"player"`
}

type Player struct {
	UUID         string `json:"uuid"`
	Achievements struct {
		Bedwars_Level int `json:"bedwars_level"`
	}
	Displayname    string   `json:"displayname"`
	FirstLogin     int64    `json:"firstLogin"`
	FriendRequests []string `json:"friendRequests"`
	Karma          int      `json:"karma"`
	KnownAliases   []string `json:"knownAliases"`
	LastLogin      int64    `json:"lastLogin"`
	NetworkExp     float32  `json:"networkExp"`
	Playername     string   `json:"playername"`
	Stats          struct {
		UUID    string `json:"uuid"`
		SkyWars struct {
			GamesPlayedSkywars int     `json:"games_played_skywars"`
			Kills              float32 `json:"kills"`
			Level              string  `json:"levelFormatted"`
			Deaths             float32 `json:"deaths"`
			BlocksPlaced       int     `json:"blocks_placed"`
			Coins              int     `json:"coins"`
			Losses             float32 `json:"losses"`
			MostKillsGame      int     `json:"most_kills_game"`
			Souls              int     `json:"souls"`
			BlocksBroken       int     `json:"blocks_broken"`
			Assists            int     `json:"assists"`
			Games              int     `json:"games"`
			ArrowsHit          int     `json:"arrows_hit"`
			FastestWin         int     `json:"fastest_win"`
			Wins               float32 `json:"wins"`
			Winstreak          int     `json:"winstreak"`
		} `json:"SkyWars"`
		Bedwars struct {
			Wins         float32 `json:"wins_bedwars"`
			Games        int     `json:"games_played_bedwars"`
			Losses       float32 `json:"losses_bedwars"`
			Final_Kills  float32 `json:"final_kills_bedwars"`
			Final_Deaths float32 `json:"final_deaths_bedwars"`
			Beds_Lost    float32 `json:"beds_lost_bedwars"`
			Beds_Broken  float32 `json:"beds_broken_bedwars"`
			Winstreak    int     `json:"winstreak"`
		} `json:"BedWars"`
	}
}

// mojang api response

type Mojang struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type Name struct {
	Name string `json:"name"`
}

// bot command

type Command struct {
	Name      string
	Args      int
	ForAdmins bool
	Trigger   func(name string, peer_id int, from_id int) (err error)
}

// status endpoint

type Status struct {
	Success bool   `json:"success"`
	UUID    string `json:"uuid"`

	Session struct {
		Online   bool   `json:"online"`
		GameType string `json:"gameType"`
		Mode     string `json:"mode"`
	}
}

type Status2 struct {
	Success bool `json:"success"`

	Player struct {
		LastLogin  int64 `json:"lastLogin"`
		LastLogout int64 `json:"lastLogout"`
	}
}

// skyblock shit god bless

type Profile struct {
	Id       string `json:"id"`
	Gamemode string `json:"game_mode"`
	CuteName string `json:"cute_name"`
	Banking  struct {
		Balance float32 `json:"balance"`
	} `json:"banking"`

	Members map[string]Member `json:"members"`
}

type ProfileReturn struct {
	Success  bool      `json:"success"`
	Profiles []Profile `json:"profiles"`
}

type Skill struct {
	Xp         float32 `json:"xp"`
	Level      int     `json:"level"`
	FloatLevel float32 `json:"floatLevel"`
	MaxLevel   int     `json:"maxLevel"`
	XpCurrent  int     `json:"xpCurrent"`
	XpForNext  int     `json:"xpForNext"`
	Progress   float32 `json:"progress"`
}

type Slayer struct {
	Claimed   int `json:"claimed_levels"`
	Xp        int `json:"xp"`
	XpForNext int `json:"xp_for_next"`
}

type Member struct {
	Purse  float32 `json:"coin_purse"`
	Skills struct {
		Mining       Skill `json:"mining"`
		Runecrafting Skill `json:"runecrafting"`
		Alchemy      Skill `json:"alchemy"`
		Taming       Skill `json:"taming"`
		Combat       Skill `json:"combat"`
		Farming      Skill `json:"farming"`
		Enchanting   Skill `json:"enchanting"`
		Fishing      Skill `json:"fishing"`
		Foraging     Skill `json:"foraging"`
		Carpentry    Skill `json:"carpentry"`
	} `json:"skills"`
	Average float32 `json:"average_skill_level"`

	Slayers struct {
		Zombie   Slayer `json:"zombie"`
		Wolf     Slayer `json:"wolf"`
		Spider   Slayer `json:"spider"`
		Enderman Slayer `json:"enderman"`
	} `json:"slayer"`

	Attributes struct {
		Damage            float32 `json:"damage"`
		Health            float32 `json:"health"`
		Defense           float32 `json:"defense"`
		EffectiveHealth   float32 `json:"effective_health"`
		Strength          float32 `json:"strength"`
		DamageIncrease    float32 `json:"damage_increase"`
		Speed             float32 `json:"speed"`
		CritChance        float32 `json:"crit_chance"`
		CritDamage        float32 `json:"crit_damage"`
		BonusAttackSpeed  float32 `json:"bonus_attack_speed"`
		Intelligence      float32 `json:"intelligence"`
		SeaCreatureChance float32 `json:"sea_creature_chance"`
		MagicFind         float32 `json:"magic_find"`
		PetLuck           float32 `json:"pet_luck"`
		TrueDefense       float32 `json:"true_defense"`
		Ferocity          float32 `json:"ferocity"`
		AbilityDamage     float32 `json:"ability_damage"`
		MiningSpeed       float32 `json:"mining_speed"`
		MiningFortune     float32 `json:"mining_fortune"`
		FarmingFortune    float32 `json:"farming_fortune"`
		ForagingFortune   float32 `json:"foraging_fortune"`
	} `json:"attributes"`
}

// Bazaar structs

type BazaarData struct {
	Success     bool
	LastUpdated int
	Products    map[string]Product
}

type Product struct {
	ID          string        `json:"product_id"`
	SellSummary []SummaryInfo `json:"sell_summary"`
	BuySummary  []SummaryInfo `json:"buy_summary"`
	QuickStatus QuickStatus   `json:"quick_status"`
}

type SummaryInfo struct {
	Amount       int
	PricePerUnit float64
	Orders       int
}

type QuickStatus struct {
	SellPrice      float64
	SellVolume     int
	SellMovingWeek int
	SellOrders     int
	BuyPrice       float64
	BuyVolume      int
	BuyMovingWeek  int
	BuyOrders      int
}

// auction data

type AuctionReturn struct {
	Success       bool
	Page          int
	TotalPages    int
	TotalAuctions int
	LastUpdated   int
	Auctions      []AuctionData
}

type AuctionCache struct {
	LastUpdated time.Time
	Auctions    []AuctionData
}

type AuctionData struct {
	ID             string `json:"uuid" gorm:"primary_key"`
	Auctioneer     string
	ProfileID      string `json:"profile_id"`
	Start          int
	End            int
	Name           string `json:"item_name"`
	Lore           string `json:"item_lore"`
	Extra          string
	Category       string
	Tier           string
	StartingBid    int `json:"starting_bid"`
	Claimed        bool
	HighestBid     int       `json:"highest_bid_amount" gorm:"index"`
	BIN            bool      `json:"bin"`
	Bids           []BidData `gorm:"foreignKey:AuctionID"`
	FinalSalePrice int
}

type BidData struct {
	AuctionID string `json:"auction_id" gorm:"primary_key;autoIncrement:false"`
	Bidder    string
	ProfileID string `json:"profile_id"`
	Amount    int
	Timestamp int `gorm:"primary_key;autoIncrement:false"`
}

type EndedAuctionReturn struct {
	Success     bool
	LastUpdated int
	Auctions    []EndedAuction
}

type EndedAuction struct {
	AuctionID     string `json:"auction_id"`
	Seller        string
	SellerProfile string
	Buyer         string
	Timestamp     int
	Price         int
	BIN           bool `json:"bin"`
}

func (auction AuctionData) GetHighestBid() BidData {
	highest := BidData{Amount: 0}
	for _, bid := range auction.Bids {
		if bid.Amount > highest.Amount {
			highest = bid
		}
	}
	return highest
}
