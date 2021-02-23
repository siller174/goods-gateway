package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/siller174/goodsGateway/pkg/gateway/config"
	"github.com/siller174/goodsGateway/pkg/logger"
	"github.com/siller174/goodsGateway/pkg/utils/rquid"
)

type Bot struct {
	cfg conf

	bot    *tgbotapi.BotAPI
	updCnf tgbotapi.UpdateConfig

	commands  map[string]func(ctx context.Context, chatID int64, text string) (*Message, error)
	callbacks map[string]func(ctx context.Context, chatID int64, query string, text string) (*Message, error)
}

type conf struct {
	adminChatId int64
	channel     string
}

func NewBot(cfg config.Telegram) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		logger.Error("Could not init telegram bot. Error %v", err)
		return nil, err
	}

	bot.Debug = cfg.DebugModeEnabled

	logger.Info("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	ourbot := &Bot{
		cfg: conf{
			adminChatId: cfg.AdminChatId,
			channel:     cfg.Channel,
		},
		bot:    bot,
		updCnf: u,
	}
	ourbot.initCallbacks()
	ourbot.initCommands()

	return ourbot, nil
}

func (bot *Bot) StartHandleMsg(ctx context.Context) {
	updates, err := bot.bot.GetUpdatesChan(bot.updCnf)

	if err != nil {
		logger.ErrorCtx(ctx, "Could not start listen telegram msg. Error %v", err)
		return
	}

	for {
		select {
		case update := <-updates:
			ctx := context.WithValue(ctx, logger.RequestLogField, rquid.CreateReqUid())

			if update.CallbackQuery != nil {
				bot.processCallback(ctx, update)
			}

			if update.Message != nil {
				bot.processUpdate(ctx, update)
			}
		case <-ctx.Done():
			logger.InfoCtx(ctx, "Context is done. Finish telegram bot")

		}
	}
}

func (bot *Bot) processUpdate(ctx context.Context, update tgbotapi.Update) {
	logger.InfoCtx(ctx, "Received message [%s][%v] %s", update.Message.From.UserName, update.Message.Chat.ID, update.Message.Text)

	// TODO ADD PROCESS UPDATE

	bot.sendMessage(update.Message.Chat.ID, update.Message.Text)
}

func (bot *Bot) processCallback(ctx context.Context, update tgbotapi.Update) {
	logger.InfoCtx(ctx, "Recieved callback %+v", update.CallbackQuery.Data)


	// TODO ADD PROCESS UPDATE
}
