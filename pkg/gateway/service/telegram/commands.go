package telegram

import "context"

func (bot *Bot) initCommands() {
	var commands = map[string]func(ctx context.Context, chatID int64, text string) (*Message, error){

	}

	bot.commands = commands
}
