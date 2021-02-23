package config

import (
	"fmt"
	"os"

	"github.com/siller174/goodsGateway/pkg/logger"
	configReader "github.com/siller174/goodsGateway/pkg/utils/config"
)

type App struct {
	DevMode bool
	Log     logger.Logger
	Server  Server
	Db      Db
	Telegram Telegram
}

func New(configPath string) App {
	var mainConfig App

	err := configReader.Read(configPath, &mainConfig)
	handleConfigError("Could not init main config", err)

	err = mainConfig.Log.Init()
	handleConfigError("Could not init logrus", err)

	logger.Info("App was configured with params: %+v", mainConfig)
	return mainConfig
}

func handleConfigError(path string, err error) {
	if err != nil {

		fmt.Printf("Could not read config from file. %s. Error: %v", path, err)
		os.Exit(-1)
	}
}
