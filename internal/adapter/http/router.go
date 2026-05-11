package httpserver

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	answerhandler "github.com/Talan-Application/api-gateway/internal/adapter/http/handler/answer"
	authhandler "github.com/Talan-Application/api-gateway/internal/adapter/http/handler/auth"
	questionhandler "github.com/Talan-Application/api-gateway/internal/adapter/http/handler/question"
	quizhandler "github.com/Talan-Application/api-gateway/internal/adapter/http/handler/quiz"
	"github.com/Talan-Application/api-gateway/internal/adapter/http/middleware"
	"github.com/Talan-Application/api-gateway/internal/usecase"
)

func NewRouter(env string, jwtSecret string, log *zap.Logger, authUC usecase.Auth, quizUC usecase.Quiz, questionUC usecase.Question, answerUC usecase.Answer) *gin.Engine {
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

	protected := v1.Group("", middleware.AuthMiddleware(jwtSecret))

	quiz := quizhandler.NewHandler(quizUC, log)
	quizGroup := protected.Group("/quizzes")
	{
		quizGroup.POST("", quiz.Create)
		quizGroup.GET("", quiz.GetAll)
		quizGroup.GET("/:id", quiz.GetByID)
		quizGroup.PUT("/:id", quiz.Update)
		quizGroup.DELETE("/:id", quiz.Delete)
	}

	question := questionhandler.NewHandler(questionUC, log)
	questionGroup := protected.Group("/questions")
	{
		questionGroup.POST("", question.Create)
		questionGroup.GET("", question.GetAll)
		questionGroup.GET("/:id", question.GetByID)
		questionGroup.PUT("/:id", question.Update)
		questionGroup.DELETE("/:id", question.Delete)
	}

	answer := answerhandler.NewHandler(answerUC, log)
	answerGroup := protected.Group("/answers")
	{
		answerGroup.POST("", answer.Create)
		answerGroup.GET("", answer.GetAll)
		answerGroup.GET("/:id", answer.GetByID)
		answerGroup.PUT("/:id", answer.Update)
		answerGroup.DELETE("/:id", answer.Delete)
	}

	return router
}
