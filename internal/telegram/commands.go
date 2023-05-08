package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *Bot) startCommand(msg *tgbotapi.Message) tgbotapi.MessageConfig {
	res := tgbotapi.NewMessage(msg.Chat.ID, msgStart)
	return res
}

func (b *Bot) setCommand(msg *tgbotapi.Message) tgbotapi.MessageConfig {
	res := tgbotapi.NewMessage(msg.Chat.ID, msgNewSet)
	b.userStates[msg.From.UserName] = SETTING
	return res
}

func (b *Bot) getCommand(msg *tgbotapi.Message) tgbotapi.MessageConfig {
	res := tgbotapi.NewMessage(msg.Chat.ID, msgNewGet)
	b.userStates[msg.From.UserName] = GETTING
	return res
}

func (b *Bot) deleteCommand(msg *tgbotapi.Message) tgbotapi.MessageConfig {
	res := tgbotapi.NewMessage(msg.Chat.ID, msgNewDel)

	b.userStates[msg.From.UserName] = DELETING
	return res
}

func (b *Bot) unknownCommand(message *tgbotapi.Message) tgbotapi.MessageConfig {
	res := tgbotapi.NewMessage(message.Chat.ID, msgUnknownCommand)
	return res
}
