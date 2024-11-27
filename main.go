package main

import (
	"api-getaway/settings"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/GOAT-prod/goatlogger"
)

func main() {
	logger := goatlogger.New(settings.AppName())
	logger.SetTag("app")

	mainCtx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	cfg, err := settings.ReadConfig()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := NewApp(mainCtx, cfg, logger)
	app.initClusterClients()
	app.initServer()
	app.initHandlers()
	app.Start()

	logger.Info(fmt.Sprintf("приложение запущено, порт: %d, конфиг: %s.json", cfg.Port, settings.GetEnv()))

	waitTerminate(mainCtx, app.Stop)
}

func waitTerminate(mainCtx context.Context, quitFn func(ctx context.Context)) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	if quitFn == nil {
		return
	}

	quitCtx, cancelQuitCtx := context.WithTimeout(mainCtx, time.Second*15)
	defer cancelQuitCtx()

	quitFn(quitCtx)
}
