package config

type Telegram struct {
	Channel          string
	Token            string
	DebugModeEnabled bool
	AdminChatId      int64
}
