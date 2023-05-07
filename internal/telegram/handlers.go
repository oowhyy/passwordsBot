package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleAny(msg *tgbotapi.Message) error {
	usr := msg.From.UserName
	// some state
	if state, ok := b.userStates[usr]; ok {
		err := b.handleState(state, msg)
		return err
	}
	// default state, some command
	if msg.IsCommand() {
		return b.handleCommand(msg)
	}
	// default
	return b.unknownCommand(msg)
}

func (b *Bot) handleCommand(msg *tgbotapi.Message) error {
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
