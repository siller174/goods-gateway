package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/siller174/goodsGateway/pkg/gateway/service/communication"
	"github.com/siller174/goodsGateway/pkg/logger"
)

func (bot *Bot) SendGoodToChannel(ctx context.Context, good communication.Good) {
	bot.sendMessageToChannel(ctx, good.ToStringForTelegram())
}

func (bot *Bot) SendGoodToChat(chatID int64, good communication.Good) {
	bot.sendMessage(chatID, good.ToStringForTelegram())
}

func (bot *Bot) SendMsgToAdmin(text string) {
	bot.sendMessage(bot.cfg.adminChatId, text)
}

func (bot *Bot) sendMessageToChannel(ctx context.Context, text string) {
	msg := tgbotapi.NewMessageToChannel(bot.cfg.channel, text)
	msg.ParseMode = "markdown"
	_, err := bot.bot.Send(msg)

	if err != nil {
		logger.ErrorCtx(ctx, "Cannot send message to channel. Message %+v. Error %v", msg, err)
	} else {
		logger.InfoCtx(ctx, "Message %+v was sent", msg)
	}
}

func (bot *Bot) sendMessage(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = "markdown"
	_, err := bot.bot.Send(msg)

	if err != nil {
		logger.Error("Cannot send message to chat. Message %+v. Error %v", msg, err)
	} else {
		logger.Info("Message %+v was sent", msg)
	}
}
