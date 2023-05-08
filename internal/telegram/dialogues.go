package telegram

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/oowhyy/passwordbot/internal/storage"
)

const maxArgLen = 200

func (b *Bot) handleState(state State, msg *tgbotapi.Message) tgbotapi.MessageConfig {
	// reset state in any case
	delete(b.userStates, msg.From.UserName)
	switch state {
	case SETTING:
		return b.setItem(msg)
	case GETTING:
		return b.getItem(msg)
	case DELETING:
		return b.delItem(msg)
	default:
		return tgbotapi.NewMessage(msg.Chat.ID, msgErrUnknown)
	}
}

// /set next step
func (b *Bot) setItem(msg *tgbotapi.Message) tgbotapi.MessageConfig {
	arr := strings.Fields(msg.Text)
	if len(arr) <= 1 {
		res := tgbotapi.NewMessage(msg.Chat.ID, msgBadSetOneWord)
		return res
	}
	// TODO check if service exists
	service := strings.Join(arr[:len(arr)-1], " ")
	password := arr[len(arr)-1]
	if len(service) > maxArgLen || len(password) > maxArgLen {
		res := tgbotapi.NewMessage(msg.Chat.ID, msgBadSetTooLong)
		return res
	}
	err := b.storage.Set(msg.From.UserName, &storage.Item{
		Service:  service,
		Password: password,
	})
	if err != nil {
		res := tgbotapi.NewMessage(msg.Chat.ID, msgErrSet)
		return res
	}
	text := fmt.Sprintf(msgOKSet, esc(password), esc(service))
	res := tgbotapi.NewMessage(msg.Chat.ID, text)
	return res
}

// /get next step
func (b *Bot) getItem(msg *tgbotapi.Message) tgbotapi.MessageConfig {
	service := strings.TrimSpace(msg.Text)
	val, err := b.storage.Get(msg.From.UserName, service)
	switch {
	case err != nil:
		res := tgbotapi.NewMessage(msg.Chat.ID, esc(msgErrGet))
		return res
	case val == nil:
		text := fmt.Sprintf(esc(msgBadGet), esc(msg.Text))
		res := tgbotapi.NewMessage(msg.Chat.ID, text)
		return res
	}
	res := tgbotapi.NewMessage(msg.Chat.ID, esc(val.Password))
	res.ReplyToMessageID = msg.MessageID
	return res
}

// /del next step
func (b *Bot) delItem(msg *tgbotapi.Message) tgbotapi.MessageConfig {
	service := strings.TrimSpace(msg.Text)
	val, err := b.storage.Delete(msg.From.UserName, service)
	if err != nil {
		res := tgbotapi.NewMessage(msg.Chat.ID, msgErrDel)
		return res
	}
	if val == 0 {
		text := fmt.Sprintf(esc(msgBadDel), esc(msg.Text))
		res := tgbotapi.NewMessage(msg.Chat.ID, text)
		return res
	}
	// val == 1
	text := fmt.Sprintf(msgOKDel, esc(msg.Text))
	res := tgbotapi.NewMessage(msg.Chat.ID, text)
	return res
}

func esc(s string) string {
	return tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, s)
}
