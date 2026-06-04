package common_subject

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Talan-Application/api-gateway/internal/model"
	"github.com/Talan-Application/api-gateway/internal/usecase"
)

type Handler struct {
	commonSubjectUC usecase.CommonSubject
	log             *zap.Logger
}

func NewHandler(commonSubjectUC usecase.CommonSubject, log *zap.Logger) *Handler {
	return &Handler{commonSubjectUC: commonSubjectUC, log: log}
}

func (h *Handler) Create(c *gin.Context) {
	var req model.CreateCommonSubjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.commonSubjectUC.CreateCommonSubject(c.Request.Context(), req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *Handler) GetByID(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	resp, err := h.commonSubjectUC.GetCommonSubject(c.Request.Context(), id)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) Lookup(c *gin.Context) {
	resp, err := h.commonSubjectUC.GetCommonSubjectsLookup(c.Request.Context())
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetAll(c *gin.Context) {
	var limit, offset *int32

	if v := c.Query("limit"); v != "" {
		n, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit"})
			return
		}
		val := int32(n)
		limit = &val
	}
	if v := c.Query("offset"); v != "" {
		n, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset"})
			return
		}
		val := int32(n)
		offset = &val
	}

	resp, err := h.commonSubjectUC.GetAllCommonSubjects(c.Request.Context(), limit, offset)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) Update(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req model.UpdateCommonSubjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.commonSubjectUC.UpdateCommonSubject(c.Request.Context(), id, req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) Delete(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	resp, err := h.commonSubjectUC.DeleteCommonSubject(c.Request.Context(), id)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func parseID(c *gin.Context) (int64, error) {
	return strconv.ParseInt(c.Param("id"), 10, 64)
}

func (h *Handler) handleError(c *gin.Context, err error) {
	st, ok := status.FromError(err)
	if !ok {
		h.log.Error("unexpected error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	switch st.Code() {
	case codes.NotFound:
		c.JSON(http.StatusNotFound, gin.H{"error": st.Message()})
	case codes.InvalidArgument:
		c.JSON(http.StatusBadRequest, gin.H{"error": st.Message()})
	case codes.AlreadyExists:
		c.JSON(http.StatusConflict, gin.H{"error": st.Message()})
	case codes.Unauthenticated:
		c.JSON(http.StatusUnauthorized, gin.H{"error": st.Message()})
	case codes.PermissionDenied:
		c.JSON(http.StatusForbidden, gin.H{"error": st.Message()})
	case codes.Unavailable:
		h.log.Error("common subject service unavailable", zap.Error(err))
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "common subject service unavailable"})
	default:
		h.log.Error("grpc error", zap.String("code", st.Code().String()), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}
