package telegram

import (
	"log"
	"time"

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
	expire     time.Duration
	messageQ   []tgbotapi.Message
}

func NewBot(api *tgbotapi.BotAPI, db storage.Storage) *Bot {
	return &Bot{
		api:        api,
		storage:    db,
		userStates: map[string]State{},
		expire:     time.Second * 10,
		messageQ:   []tgbotapi.Message{},
	}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.api.Self.UserName)
	go func() {
		for {
			time.Sleep(time.Second * 5)
			b.clearChat()
		}
	}()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := b.api.GetUpdatesChan(u)
	for update := range updates {
		log.Println("HERE")
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		res := b.HandleAny(update.Message)
		res.ParseMode = "MarkdownV2"
		sent, err := b.api.Send(res)
		b.messageQ = append(b.messageQ, sent)
		if err != nil {
			log.Println("send reply failed")
		}
	}
	return nil
}

func (b *Bot) clearChat() {
	for len(b.messageQ) > 0 {
		m := b.messageQ[0]
		t := time.Unix(int64(m.Date), 0)
		if time.Since(t) > b.expire {
			b.messageQ = b.messageQ[1:]
			delete := tgbotapi.NewDeleteMessage(m.Chat.ID, m.MessageID)
			b.api.Request(delete)
		} else {
			break
		}
	}
}
