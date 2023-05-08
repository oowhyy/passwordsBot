package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/oowhyy/passwordbot/internal/storage"
)

type State int

const (
	DEFAULT State = iota
	SETTING
	GETTING
	DELETING
)

type Bot struct {
	api        *tgbotapi.BotAPI
	storage    storage.Storage
	userStates map[string]State
}

func NewBot(api *tgbotapi.BotAPI, db storage.Storage) *Bot {
	return &Bot{
		api:        api,
		storage:    db,
		userStates: map[string]State{},
	}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.api.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 15
	updates := b.api.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		res := b.HandleAny(update.Message)
		_, err := b.api.Send(res)
		if err != nil {
			log.Println("send reply failed")
		}
	}
	return nil
}
