package httpserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/Talan-Application/api-gateway/internal/config"
)

type Server struct {
	srv *http.Server
	log *zap.Logger
}

func NewServer(cfg config.HTTPServerConfig, router *gin.Engine, log *zap.Logger) *Server {
	return &Server{
		srv: &http.Server{
			Addr:           fmt.Sprintf(":%d", cfg.Port),
			Handler:        router,
			ReadTimeout:    cfg.ReadTimeout,
			WriteTimeout:   cfg.WriteTimeout,
			IdleTimeout:    cfg.IdleTimeout,
			MaxHeaderBytes: cfg.MaxHeaderBytes,
		},
		log: log,
	}
}

func (s *Server) Run() error {
	s.log.Info("HTTP server started", zap.String("addr", s.srv.Addr))
	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
