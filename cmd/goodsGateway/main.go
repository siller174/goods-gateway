package main

import (
	"context"
	"github.com/siller174/goodsGateway/pkg/gateway/api"
	"github.com/siller174/goodsGateway/pkg/gateway/config"
	"github.com/siller174/goodsGateway/pkg/logger"
	"github.com/spf13/pflag"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-ch
		logger.Info("Handle signal %v. Exiting", sig)
		cancel()
		os.Exit(0)
	}()

	configPath := pflag.String("config-path", "./goodsGateway.properties", "Path to config file")
	pflag.Parse()
	appConfig := config.New(*configPath)

	server, err := api.New(ctx, appConfig)
	if err != nil {
		logger.Fatal("cannot start app. Init api error %v", err)
	}
	logger.Fatal("App was close", server.ListenAndServe())
}
