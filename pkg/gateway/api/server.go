package api

import (
	"context"
	"github.com/siller174/goodsGateway/pkg/gateway/repository"
	"github.com/siller174/goodsGateway/pkg/gateway/service/communication"
	"github.com/siller174/goodsGateway/pkg/gateway/service/telegram"
	"github.com/siller174/goodsGateway/pkg/logger"
	"net/http"

	"github.com/siller174/goodsGateway/pkg/gateway/api/handler"
	errHandler "github.com/siller174/goodsGateway/pkg/gateway/api/http/errors/handler"
	"github.com/siller174/goodsGateway/pkg/gateway/api/middleware"

	"github.com/gorilla/mux"
	manageApi "github.com/siller174/goodsGateway/pkg/gateway/api/manage"
	"github.com/siller174/goodsGateway/pkg/gateway/config"
	"github.com/siller174/goodsGateway/pkg/gateway/service/manage"
)

func New(ctx context.Context, appConfig config.App) (*http.Server, error) {
	router, err := initRouters(ctx, appConfig)

	if err != nil {
		return nil, err
	}

	srv := http.Server{
		Handler:      router,
		Addr:         appConfig.Server.Port,
		WriteTimeout: appConfig.Server.WriteTimeout,
		ReadTimeout:  appConfig.Server.ReadTimeout,
	}
	return &srv, nil
}

func initRouters(ctx context.Context, config config.App) (*mux.Router, error) {
	// Init database
	mysql, err := repository.NewMysql(ctx, config.Db)

	if err != nil {
		logger.Error("Cannot create db connect. Error %v", err)
		return nil, err
	}

	// Init telegram bot
	bot, err := telegram.NewBot(config.Telegram)
	if err != nil {
		logger.Fatal("Cannot init bot. Error %v", err)
	}
	go bot.StartHandleMsg(ctx)

	// Init media sender

	commutator := communication.NewCommutator(bot)

	manager := manage.NewManage(mysql, commutator, bot)

	// Init API
	healthApi := manageApi.NewHealthApi()
	errorHandler := errHandler.NewErrorHandler(config.DevMode)
	middleWare := middleware.NewMiddleWare(errorHandler)

	meetingApi := handler.NewHandler(manager)

	router := mux.NewRouter()
	router.HandleFunc(manageApi.HealthRoute, healthApi.Handle()).Methods(http.MethodGet)
	router.HandleFunc(handler.RouteGetCatalogs, middleWare.Handle(meetingApi.GetCatalogs)).Methods(http.MethodGet)
	router.HandleFunc(handler.RouteSaveGoods, middleWare.Handle(meetingApi.SaveGoods)).Methods(http.MethodPost)
	router.HandleFunc(handler.RouteSaveLogs, middleWare.Handle(meetingApi.SaveLogs)).Methods(http.MethodPost)
	router.Use(middleWare.RecoverPanic, middleWare.ContextRequest)

	return router, nil
}
