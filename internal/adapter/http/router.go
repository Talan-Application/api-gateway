package httpserver

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	answerhandler "github.com/Talan-Application/api-gateway/internal/adapter/http/handler/answer"
	authhandler "github.com/Talan-Application/api-gateway/internal/adapter/http/handler/auth"
	questionhandler "github.com/Talan-Application/api-gateway/internal/adapter/http/handler/question"
	quizhandler "github.com/Talan-Application/api-gateway/internal/adapter/http/handler/quiz"
	commonsubjecthandler "github.com/Talan-Application/api-gateway/internal/adapter/http/handler/common_subject"
	"github.com/Talan-Application/api-gateway/internal/adapter/http/middleware"
	"github.com/Talan-Application/api-gateway/internal/usecase"
)

func NewRouter(env string, jwtSecret string, log *zap.Logger, authUC usecase.Auth, quizUC usecase.Quiz, questionUC usecase.Question, answerUC usecase.Answer, commonSubjectUC usecase.CommonSubject) *gin.Engine {
	if env != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.CorsMiddleware())
	router.Use(middleware.LocaleMiddleware())

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
	staffOnly := middleware.RoleMiddleware("curator", "teacher", "admin")

	quiz := quizhandler.NewHandler(quizUC, log)
	question := questionhandler.NewHandler(questionUC, log)
	answer := answerhandler.NewHandler(answerUC, log)

	quizGroup := protected.Group("/quizzes")
	{
		quizGroup.GET("/:id/take", quiz.TakeQuiz)
		quizGroup.POST("/:id/submit", quiz.SubmitQuiz)
		quizGroup.GET("/:id/results", quiz.GetResults)

		quizCRUD := quizGroup.Group("", staffOnly)
		{
			quizCRUD.POST("", quiz.Create)
			quizCRUD.GET("", quiz.GetAll)
			quizCRUD.GET("/:id", quiz.GetByID)
			quizCRUD.PUT("/:id", quiz.Update)
			quizCRUD.DELETE("/:id", quiz.Delete)
			quizCRUD.GET("/:id/questions", question.GetAll)
		}
	}

	questionGroup := protected.Group("/questions", staffOnly)
	{
		questionGroup.POST("", question.Create)
		questionGroup.GET("/:id", question.GetByID)
		questionGroup.PUT("/:id", question.Update)
		questionGroup.DELETE("/:id", question.Delete)
		questionGroup.GET("/:id/answers", answer.GetAll)
	}

	answerGroup := protected.Group("/answers", staffOnly)
	{
		answerGroup.POST("", answer.Create)
		answerGroup.GET("/:id", answer.GetByID)
		answerGroup.PUT("/:id", answer.Update)
		answerGroup.DELETE("/:id", answer.Delete)
	}

	commonSubject := commonsubjecthandler.NewHandler(commonSubjectUC, log)
	commonSubjectGroup := protected.Group("/common-subjects")
	{
		commonSubjectGroup.GET("", commonSubject.GetAll)
		commonSubjectGroup.GET("/:id", commonSubject.GetByID)

		commonSubjectStaff := commonSubjectGroup.Group("", staffOnly)
		{
			commonSubjectStaff.POST("", commonSubject.Create)
			commonSubjectStaff.PUT("/:id", commonSubject.Update)
			commonSubjectStaff.DELETE("/:id", commonSubject.Delete)
		}
	}

	return router
}
