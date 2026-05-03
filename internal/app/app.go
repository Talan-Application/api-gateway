package app

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"

	authgrpc "github.com/Talan-Application/api-gateway/internal/adapter/grpc/auth"
	httpserver "github.com/Talan-Application/api-gateway/internal/adapter/http"
	"github.com/Talan-Application/api-gateway/internal/config"
	authusecase "github.com/Talan-Application/api-gateway/internal/usecase/auth"
	"github.com/Talan-Application/api-gateway/pkg/grpcconn"
)

const shutdownTimeout = 5 * time.Second

type App struct {
	httpServer *httpserver.Server
	log        *zap.Logger
}

func New(cfg *config.Config, log *zap.Logger) (*App, error) {
	authConn, err := grpcconn.New(cfg.Services.Auth.Address)
	if err != nil {
		return nil, fmt.Errorf("auth grpc connection: %w", err)
	}

	authClient := authgrpc.NewClient(authConn)
	authUC := authusecase.New(authClient)

	router := httpserver.NewRouter(cfg.App.Env, log, authUC)
	srv := httpserver.NewServer(cfg.Server.HTTPServer, router, log)

	return &App{httpServer: srv, log: log}, nil
}

func (a *App) Run() error {
	return a.httpServer.Run()
}

func (a *App) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := a.httpServer.Shutdown(ctx); err != nil {
		a.log.Error("http server shutdown", zap.Error(err))
	}
}
