package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Talan-Application/api-gateway/internal/model"
	"github.com/Talan-Application/api-gateway/internal/usecase"
)

type Handler struct {
	authUC usecase.Auth
	log    *zap.Logger
}

func NewHandler(authUC usecase.Auth, log *zap.Logger) *Handler {
	return &Handler{authUC: authUC, log: log}
}

func (h *Handler) Register(c *gin.Context) {
	var req model.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.authUC.Register(c.Request.Context(), req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *Handler) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.authUC.Login(c.Request.Context(), req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) RefreshToken(c *gin.Context) {
	var req model.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.authUC.RefreshToken(c.Request.Context(), req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) handleError(c *gin.Context, err error) {
	st, ok := status.FromError(err)
	if !ok {
		h.log.Error("unexpected error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	switch st.Code() {
	case codes.AlreadyExists:
		c.JSON(http.StatusConflict, gin.H{"error": st.Message()})
	case codes.NotFound:
		c.JSON(http.StatusNotFound, gin.H{"error": st.Message()})
	case codes.Unauthenticated:
		c.JSON(http.StatusUnauthorized, gin.H{"error": st.Message()})
	case codes.InvalidArgument:
		c.JSON(http.StatusBadRequest, gin.H{"error": st.Message()})
	case codes.Unavailable:
		h.log.Error("auth service unavailable", zap.Error(err))
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "auth service unavailable"})
	default:
		h.log.Error("grpc error", zap.String("code", st.Code().String()), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}
