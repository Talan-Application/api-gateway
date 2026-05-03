package main

import (
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/Talan-Application/api-gateway/internal/app"
	"github.com/Talan-Application/api-gateway/internal/config"
	"github.com/Talan-Application/api-gateway/pkg/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic("load config: " + err.Error())
	}

	log := logger.New(cfg.App.Env)
	defer log.Sync()

	a, err := app.New(cfg, log)
	if err != nil {
		log.Fatal("init app", zap.Error(err))
	}

	go func() {
		if err := a.Run(); err != nil {
			log.Error("app run", zap.Error(err))
		}
	}()

	log.Info("api-gateway started", zap.String("env", cfg.App.Env))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	a.Stop()
	log.Info("api-gateway stopped")
}
