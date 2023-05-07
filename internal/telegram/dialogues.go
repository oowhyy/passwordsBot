package telegram

import (
	"errors"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/oowhyy/passwordbot/internal/storage"
)

func (b *Bot) handleState(state State, msg *tgbotapi.Message) error {
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
		return errors.New("unknown state")
	}
}

// /set next step
func (b *Bot) setItem(msg *tgbotapi.Message) error {
	arr := strings.Fields(msg.Text)
	if len(arr) <= 1 {
		res := tgbotapi.NewMessage(msg.Chat.ID, msgBadSet)
		b.api.Send(res)
		return nil
	}
	// TODO check if service exists
	service := strings.Join(arr[:len(arr)-1], " ")
	password := arr[len(arr)-1]
	err := b.storage.Set(msg.From.UserName, &storage.Item{
		Service:  service,
		Password: password,
	})
	if err != nil {
		return err
	}
	text := fmt.Sprintf(msgOKSet, password, service)
	res := tgbotapi.NewMessage(msg.Chat.ID, text)
	_, err = b.api.Send(res)
	return err
}

// /get next step
func (b *Bot) getItem(msg *tgbotapi.Message) error {
	service := strings.TrimSpace(msg.Text)
	val, err := b.storage.Get(msg.From.UserName, service)
	switch {
	case err != nil:
		return err
	case val == nil:
		text := fmt.Sprintf(msgBadGet, msg.Text)
		res := tgbotapi.NewMessage(msg.Chat.ID, text)
		_, err = b.api.Send(res)
		return err
	}
	res := tgbotapi.NewMessage(msg.Chat.ID, val.Password)
	res.ReplyToMessageID = msg.MessageID
	_, err = b.api.Send(res)
	return err
}

// /del next step
func (b *Bot) delItem(msg *tgbotapi.Message) error {
	service := strings.TrimSpace(msg.Text)
	val, err := b.storage.Delete(msg.From.UserName, service)
	if err != nil {
		return err
	}
	if val == 0 {
		text := fmt.Sprintf(msgBadDel, msg.Text)
		res := tgbotapi.NewMessage(msg.Chat.ID, text)
		_, err = b.api.Send(res)
		return err
	}
	// val == 1
	text := fmt.Sprintf(msgOKDel, msg.Text)
	res := tgbotapi.NewMessage(msg.Chat.ID, text)
	_, err = b.api.Send(res)
	return err
}
