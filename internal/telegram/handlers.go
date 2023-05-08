package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) HandleAny(msg *tgbotapi.Message) tgbotapi.MessageConfig {
	b.messageQ = append(b.messageQ, *msg)
	usr := msg.From.UserName
	// some state
	if state, ok := b.userStates[usr]; ok {
		return b.handleState(state, msg)
	}
	// default state, some command
	if msg.IsCommand() {
		return b.handleCommand(msg)
	}
	// default
	return b.unknownCommand(msg)
}

func (b *Bot) handleCommand(msg *tgbotapi.Message) tgbotapi.MessageConfig {
	switch msg.Command() {
	case commandStart:
		return b.startCommand(msg)
	case commandSet:
		return b.setCommand(msg)
	case commandGet:
		return b.getCommand(msg)
	case commandDelete:
		return b.deleteCommand(msg)
	default:
		return b.unknownCommand(msg)
	}
}
