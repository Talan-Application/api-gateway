package app

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"

	answergrpc "github.com/Talan-Application/api-gateway/internal/adapter/grpc/answer"
	authgrpc "github.com/Talan-Application/api-gateway/internal/adapter/grpc/auth"
	questiongrpc "github.com/Talan-Application/api-gateway/internal/adapter/grpc/question"
	quizgrpc "github.com/Talan-Application/api-gateway/internal/adapter/grpc/quiz"
	httpserver "github.com/Talan-Application/api-gateway/internal/adapter/http"
	"github.com/Talan-Application/api-gateway/internal/config"
	answerusecase "github.com/Talan-Application/api-gateway/internal/usecase/answer"
	authusecase "github.com/Talan-Application/api-gateway/internal/usecase/auth"
	questionusecase "github.com/Talan-Application/api-gateway/internal/usecase/question"
	quizusecase "github.com/Talan-Application/api-gateway/internal/usecase/quiz"
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

	quizConn, err := grpcconn.New(cfg.Services.Quiz.Address)
	if err != nil {
		return nil, fmt.Errorf("quiz grpc connection: %w", err)
	}

	authClient := authgrpc.NewClient(authConn)
	authUC := authusecase.New(authClient)

	quizClient := quizgrpc.NewClient(quizConn)
	quizUC := quizusecase.New(quizClient)

	questionClient := questiongrpc.NewClient(quizConn)
	questionUC := questionusecase.New(questionClient)

	answerClient := answergrpc.NewClient(quizConn)
	answerUC := answerusecase.New(answerClient)

	router := httpserver.NewRouter(cfg.App.Env, cfg.JWT.SecretKey, log, authUC, quizUC, questionUC, answerUC)
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
