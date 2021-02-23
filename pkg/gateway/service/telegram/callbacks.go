package telegram

import "context"

const (
	delimiter = "-"
)

func (bot *Bot) initCallbacks() {
	var callbacks = map[string]func(ctx context.Context, chatID int64, query string, text string) (*Message, error){

	}
	bot.callbacks = callbacks
}

func getCallbackCmd(query string) string {
	tmp := []rune(query)
	res := make([]rune, 0)

	for _, r := range tmp {
		if string(r) == delimiter {
			break
		}
		res = append(res, r)
	}
	return string(res)
}
