package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *Bot) startCommand(msg *tgbotapi.Message) error {
	text := msgStart
	res := tgbotapi.NewMessage(msg.Chat.ID, text)
	_, err := b.api.Send(res)
	return err
}

func (b *Bot) setCommand(msg *tgbotapi.Message) error {
	res := tgbotapi.NewMessage(msg.Chat.ID, msgNewSet)
	_, err := b.api.Send(res)
	if err != nil {
		return err
	}
	b.userStates[msg.From.UserName] = SETTING
	return nil
}

func (b *Bot) getCommand(msg *tgbotapi.Message) error {
	res := tgbotapi.NewMessage(msg.Chat.ID, msgNewGet)
	_, err := b.api.Send(res)
	if err != nil {
		return err
	}
	b.userStates[msg.From.UserName] = GETTING
	return nil
}

func (b *Bot) deleteCommand(msg *tgbotapi.Message) error {
	res := tgbotapi.NewMessage(msg.Chat.ID, msgNewDel)
	_, err := b.api.Send(res)
	if err != nil {
		return err
	}
	b.userStates[msg.From.UserName] = DELETING
	return nil
}

func (b *Bot) unknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, msgUnknownCommand)
	_, err := b.api.Send(msg)
	return err
}
