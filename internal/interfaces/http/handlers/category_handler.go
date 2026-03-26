package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"pesa-mind/internal/domain/category"
	"pesa-mind/internal/interfaces/http/dto"
)

type CategoryHandler struct {
	Service *category.Service
}

func NewCategoryHandler(s *category.Service) *CategoryHandler {
	return &CategoryHandler{Service: s}
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userIDStr, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID, err := uuid.Parse(userIDStr.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	var parentID *uuid.UUID
	if req.ParentID != nil {
		id, err := uuid.Parse(*req.ParentID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid parent_id"})
			return
		}
		parentID = &id
	}
	category, err := h.Service.Create(userID, req.Name, req.Type, parentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var parentIDStr *string
	if category.ParentID != nil {
		str := category.ParentID.String()
		parentIDStr = &str
	}
	c.JSON(http.StatusCreated, dto.CategoryResponse{
		ID:       category.ID.String(),
		UserID:   category.UserID.String(),
		Name:     category.Name,
		Type:     category.Type,
		ParentID: parentIDStr,
	})
}

func (h *CategoryHandler) List(c *gin.Context) {
	userIDStr, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID, err := uuid.Parse(userIDStr.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	categories, err := h.Service.GetByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var resp []dto.CategoryResponse
	for _, cat := range categories {
		var parentIDStr *string
		if cat.ParentID != nil {
			str := cat.ParentID.String()
			parentIDStr = &str
		}
		resp = append(resp, dto.CategoryResponse{
			ID:       cat.ID.String(),
			UserID:   cat.UserID.String(),
			Name:     cat.Name,
			Type:     cat.Type,
			ParentID: parentIDStr,
		})
	}
	c.JSON(http.StatusOK, resp)
}
