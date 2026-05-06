package httpserver

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	authhandler "github.com/Talan-Application/api-gateway/internal/adapter/http/handler/auth"
	"github.com/Talan-Application/api-gateway/internal/adapter/http/middleware"
	"github.com/Talan-Application/api-gateway/internal/usecase"
)

func NewRouter(env string, log *zap.Logger, authUC usecase.Auth) *gin.Engine {
	if env != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.CorsMiddleware())

	v1 := router.Group("/api/v1")

	auth := authhandler.NewHandler(authUC, log)
	authGroup := v1.Group("/auth")
	{
		authGroup.POST("/register", auth.Register)
		authGroup.POST("/login", auth.Login)
		authGroup.POST("/verify-email", auth.VerifyEmail)
		authGroup.POST("/verify-login", auth.VerifyLoginCode)
		authGroup.POST("/refresh", auth.RefreshToken)
	}

	return router
}
