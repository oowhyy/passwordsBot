package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/oowhyy/passwordbot/config"
	"github.com/oowhyy/passwordbot/internal/storage/redis"
	"github.com/oowhyy/passwordbot/internal/telegram"
)

func main() {
	cfg := config.GetConfig()

	// database setup
	addr := fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)
	rdb := redis.NewRedisStorage(addr, cfg.Redis.Password, cfg.Redis.DB, cfg.Redis.ExpireMinutes)
	log.Println("connected to redis")

	// bot setup
	botApi, err := tgbotapi.NewBotAPI(cfg.Telegram.Token)
	if err != nil {
		log.Fatal(err)
	}
	botApi.Debug = true
	bot := telegram.NewBot(botApi, rdb)
	log.Println("created bot")
	if true {
		err := bot.Start()
		if err != nil {
			log.Fatal(err)
		}
	}
}

// {"ok":true,"result":[{"update_id":900651360,
// "message":{"message_id":11,"from":{"id":5717330363,"is_bot":false,
// "first_name":"\u0412\u0438\u043a\u0442\u043e\u0440","username":"oowhyy","language_code":"ru"},
// "chat":{"id":5717330363,"first_name":"\u0412\u0438\u043a\u0442\u043e\u0440",
// "username":"oowhyy","type":"private"},"date":1683450004,"text":"123"}}]}
