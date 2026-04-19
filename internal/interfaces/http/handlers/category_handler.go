package handlers

import (
	"net/http"
	"pesa-mind/internal/domain/category"
	userDomain "pesa-mind/internal/domain/user"
	"pesa-mind/internal/interfaces/http/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CategoryHandler struct {
	Service *category.Service
}

func NewCategoryHandler(s *category.Service) *CategoryHandler {
	return &CategoryHandler{Service: s}
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var req dto.CreateChannelDetailsRequest
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
	// Create a user reference for the channel details
	user := &userDomain.User{}
	user.ID = userID

	channelType := userDomain.ChannelType(req.ChannelType)

	channelDetails, err := h.Service.Create(user, req.Name, req.Description, &channelType, req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.ChannelDetailsResponse{
		ID:          channelDetails.ID.String(),
		UserID:      channelDetails.User.ID.String(),
		Name:        channelDetails.Name,
		Description: channelDetails.Description,
		ChannelType: string(*channelDetails.ChannelType),
		Status:      channelDetails.Status,
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
	channelDetails, err := h.Service.GetByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var resp []dto.ChannelDetailsResponse
	for _, cd := range channelDetails {
		resp = append(resp, dto.ChannelDetailsResponse{
			ID:          cd.ID.String(),
			UserID:      cd.User.ID.String(),
			Name:        cd.Name,
			Description: cd.Description,
			ChannelType: string(*cd.ChannelType),
			Status:      cd.Status,
		})
	}
	c.JSON(http.StatusOK, resp)
}

func (h *CategoryHandler) GetByChannelType(c *gin.Context) {
	channelTypeStr := c.Query("channel_type")
	if channelTypeStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "channel_type query parameter is required"})
		return
	}
	channelType := userDomain.ChannelType(channelTypeStr)
	channelDetails, err := h.Service.GetByChannelType(channelType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var resp []dto.ChannelDetailsResponse
	for _, cd := range channelDetails {
		resp = append(resp, dto.ChannelDetailsResponse{
			ID:          cd.ID.String(),
			UserID:      cd.User.ID.String(),
			Name:        cd.Name,
			Description: cd.Description,
			ChannelType: string(*cd.ChannelType),
			Status:      cd.Status,
		})
	}
	c.JSON(http.StatusOK, resp)
}

func (h *CategoryHandler) GetByStatus(c *gin.Context) {
	statusStr := c.Query("status")
	if statusStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status query parameter is required"})
		return
	}
	status := statusStr == "true"
	channelDetails, err := h.Service.GetByStatus(status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var resp []dto.ChannelDetailsResponse
	for _, cd := range channelDetails {
		resp = append(resp, dto.ChannelDetailsResponse{
			ID:          cd.ID.String(),
			UserID:      cd.User.ID.String(),
			Name:        cd.Name,
			Description: cd.Description,
			ChannelType: string(*cd.ChannelType),
			Status:      cd.Status,
		})
	}
	c.JSON(http.StatusOK, resp)
}

func (h *CategoryHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var req dto.UpdateChannelDetailsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	channelDetails := &category.ChannelDetails{
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
	}
	channelDetails.ID = id
	if err := h.Service.Update(channelDetails); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "updated successfully"})
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.Service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
}
