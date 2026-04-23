package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"pesa-mind/internal/domain/budget"
	"pesa-mind/internal/infrastructure/utils"
	"pesa-mind/internal/interfaces/http/dto"
	"strconv"
	"time"
)

type BudgetHandler struct {
	Service *budget.Service
}

func NewBudgetHandler(s *budget.Service) *BudgetHandler {
	return &BudgetHandler{Service: s}
}

// Budget endpoints
func (h *BudgetHandler) Create(c *gin.Context) {
	var req dto.CreateBudgetRequest
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
	b, err := h.Service.Create(userID, req.Name, req.Amount, req.Period, req.StartDate, req.EndDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.BudgetResponse{
		ID:        b.ID.String(),
		UserID:    b.UserID.String(),
		Name:      b.Name,
		Amount:    b.Amount,
		Period:    b.Period,
		StartDate: b.StartDate.Unix(),
		EndDate:   b.EndDate.Unix(),
	})
}

func (h *BudgetHandler) List(c *gin.Context) {
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
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	budgets, err := h.Service.List(userID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var resp []dto.BudgetResponse
	for _, b := range budgets {
		resp = append(resp, dto.BudgetResponse{
			ID:        b.ID.String(),
			UserID:    b.UserID.String(),
			Name:      b.Name,
			Amount:    b.Amount,
			Period:    b.Period,
			StartDate: b.StartDate.Unix(),
			EndDate:   b.EndDate.Unix(),
		})
	}
	c.JSON(http.StatusOK, resp)
}

func (h *BudgetHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	budgetID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid budget id"})
		return
	}
	b, err := h.Service.GetByID(budgetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if b == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "budget not found"})
		return
	}
	c.JSON(http.StatusOK, dto.BudgetResponse{
		ID:        b.ID.String(),
		UserID:    b.UserID.String(),
		Name:      b.Name,
		Amount:    b.Amount,
		Period:    b.Period,
		StartDate: b.StartDate.Unix(),
		EndDate:   b.EndDate.Unix(),
	})
}

func (h *BudgetHandler) Update(c *gin.Context) {
	id := c.Param("id")
	budgetID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid budget id"})
		return
	}
	var req dto.UpdateBudgetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	b, err := h.Service.GetByID(budgetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if b == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "budget not found"})
		return
	}
	if req.Name != "" {
		b.Name = req.Name
	}
	if req.Amount > 0 {
		b.Amount = req.Amount
	}
	if req.Period != "" {
		b.Period = req.Period
	}
	if req.StartDate > 0 {
		b.StartDate = time.Unix(req.StartDate, 0).UTC()
	}
	if req.EndDate > 0 {
		b.EndDate = time.Unix(req.EndDate, 0).UTC()
	}
	if err := h.Service.Update(b); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.BudgetResponse{
		ID:        b.ID.String(),
		UserID:    b.UserID.String(),
		Name:      b.Name,
		Amount:    b.Amount,
		Period:    b.Period,
		StartDate: b.StartDate.Unix(),
		EndDate:   b.EndDate.Unix(),
	})
}

func (h *BudgetHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	budgetID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid budget id"})
		return
	}
	if err := h.Service.Delete(budgetID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// MonthlyBudget endpoints
func (h *BudgetHandler) CreateMonthlyBudget(c *gin.Context) {
	var req dto.CreateMonthlyBudgetRequest
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

	// Convert request transactions to domain transactions
	transactions := make([]budget.BudgetTransaction, len(req.Transactions))
	for i, bt := range req.Transactions {
		transactions[i] = budget.BudgetTransaction{
			Name:   bt.Name,
			Amount: bt.Amount,
			Type:   utils.TransactionType(bt.Type),
		}
	}

	mb, err := h.Service.CreateMonthlyBudget(userID, req.YearlyBudgetID, time.Month(req.Month), req.Year, transactions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, h.buildMonthlyBudgetResponse(mb))
}

func (h *BudgetHandler) GetMonthlyBudgetByID(c *gin.Context) {
	id := c.Param("id")
	budgetID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid budget id"})
		return
	}
	mb, err := h.Service.GetMonthlyBudgetByID(budgetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if mb == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "monthly budget not found"})
		return
	}
	c.JSON(http.StatusOK, h.buildMonthlyBudgetResponse(mb))
}

func (h *BudgetHandler) GetMonthlyBudgetsByYearlyBudgetID(c *gin.Context) {
	yearlyBudgetID := c.Param("yearly_budget_id")
	id, err := uuid.Parse(yearlyBudgetID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid yearly budget id"})
		return
	}
	mbs, err := h.Service.GetMonthlyBudgetsByYearlyBudgetID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var resp []dto.MonthlyBudgetResponse
	for _, mb := range mbs {
		resp = append(resp, h.buildMonthlyBudgetResponse(&mb))
	}
	c.JSON(http.StatusOK, resp)
}

func (h *BudgetHandler) GetMonthlyBudgetByUserIDAndMonthYear(c *gin.Context) {
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
	month, _ := strconv.Atoi(c.Query("month"))
	year, _ := strconv.ParseInt(c.Query("year"), 10, 64)

	if month < 1 || month > 12 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid month"})
		return
	}

	mb, err := h.Service.GetMonthlyBudgetByUserIDAndMonthYear(userID, time.Month(month), year)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if mb == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "monthly budget not found"})
		return
	}
	c.JSON(http.StatusOK, h.buildMonthlyBudgetResponse(mb))
}

func (h *BudgetHandler) ListMonthlyBudgets(c *gin.Context) {
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
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	mbs, err := h.Service.GetMonthlyBudgetsByUserID(userID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var resp []dto.MonthlyBudgetResponse
	for _, mb := range mbs {
		resp = append(resp, h.buildMonthlyBudgetResponse(&mb))
	}
	c.JSON(http.StatusOK, resp)
}

func (h *BudgetHandler) UpdateMonthlyBudget(c *gin.Context) {
	id := c.Param("id")
	budgetID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid budget id"})
		return
	}
	var req dto.CreateMonthlyBudgetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mb, err := h.Service.GetMonthlyBudgetByID(budgetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if mb == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "monthly budget not found"})
		return
	}

	// Update transactions
	transactions := make([]budget.BudgetTransaction, len(req.Transactions))
	for i, bt := range req.Transactions {
		transactions[i] = budget.BudgetTransaction{
			Name:   bt.Name,
			Amount: bt.Amount,
			Type:   utils.TransactionType(bt.Type),
		}
	}
	mb.BudgetTransactions = transactions

	if err := h.Service.UpdateMonthlyBudget(mb); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, h.buildMonthlyBudgetResponse(mb))
}

func (h *BudgetHandler) DeleteMonthlyBudget(c *gin.Context) {
	id := c.Param("id")
	budgetID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid budget id"})
		return
	}
	if err := h.Service.DeleteMonthlyBudget(budgetID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// YearlyBudget endpoints
func (h *BudgetHandler) CreateYearlyBudget(c *gin.Context) {
	var req dto.CreateYearlyBudgetRequest
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

	// Convert request transactions to domain transactions
	transactions := make([]budget.BudgetTransaction, len(req.Transactions))
	for i, bt := range req.Transactions {
		transactions[i] = budget.BudgetTransaction{
			Name:   bt.Name,
			Amount: bt.Amount,
			Type:   utils.TransactionType(bt.Type),
		}
	}

	yb, err := h.Service.CreateYearlyBudget(userID, req.Year, transactions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, h.buildYearlyBudgetResponse(yb))
}

func (h *BudgetHandler) GetYearlyBudgetByID(c *gin.Context) {
	id := c.Param("id")
	budgetID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid budget id"})
		return
	}
	yb, err := h.Service.GetYearlyBudgetByID(budgetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if yb == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "yearly budget not found"})
		return
	}
	c.JSON(http.StatusOK, h.buildYearlyBudgetResponse(yb))
}

func (h *BudgetHandler) ListYearlyBudgets(c *gin.Context) {
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
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	ybs, err := h.Service.GetYearlyBudgetsByUserID(userID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var resp []dto.YearlyBudgetResponse
	for _, yb := range ybs {
		resp = append(resp, h.buildYearlyBudgetResponse(&yb))
	}
	c.JSON(http.StatusOK, resp)
}

func (h *BudgetHandler) UpdateYearlyBudget(c *gin.Context) {
	id := c.Param("id")
	budgetID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid budget id"})
		return
	}
	var req dto.CreateYearlyBudgetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	yb, err := h.Service.GetYearlyBudgetByID(budgetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if yb == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "yearly budget not found"})
		return
	}

	// Update transactions
	transactions := make([]budget.BudgetTransaction, len(req.Transactions))
	for i, bt := range req.Transactions {
		transactions[i] = budget.BudgetTransaction{
			Name:   bt.Name,
			Amount: bt.Amount,
			Type:   utils.TransactionType(bt.Type),
		}
	}
	yb.BudgetTransactions = transactions

	if err := h.Service.UpdateYearlyBudget(yb); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, h.buildYearlyBudgetResponse(yb))
}

func (h *BudgetHandler) DeleteYearlyBudget(c *gin.Context) {
	id := c.Param("id")
	budgetID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid budget id"})
		return
	}
	if err := h.Service.DeleteYearlyBudget(budgetID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// Helper methods
func (h *BudgetHandler) buildMonthlyBudgetResponse(mb *budget.MonthlyBudget) dto.MonthlyBudgetResponse {
	transactions := make([]dto.BudgetTransactionResponse, len(mb.BudgetTransactions))
	for i, bt := range mb.BudgetTransactions {
		transactions[i] = dto.BudgetTransactionResponse{
			ID:        bt.ID.String(),
			Name:      bt.Name,
			Amount:    bt.Amount,
			Type:      string(bt.Type),
			CreatedAt: bt.CreatedAt.Unix(),
		}
	}
	return dto.MonthlyBudgetResponse{
		ID:                mb.ID.String(),
		UserID:            mb.UserID.String(),
		YearlyBudgetID:    mb.YearlyBudgetID.String(),
		Month:             int(mb.Month),
		Year:              mb.Year,
		TotalExpenditures: mb.TotalExpenditures,
		TotalIncome:       mb.TotalIncome,
		TotalSavings:      mb.TotalSavings,
		TotalTransactions: mb.TotalTransactions,
		Transactions:      transactions,
		CreatedAt:         mb.CreatedAt.Unix(),
		UpdatedAt:         mb.UpdatedAt.Unix(),
	}
}

func (h *BudgetHandler) buildYearlyBudgetResponse(yb *budget.YearlyBudget) dto.YearlyBudgetResponse {
	transactions := make([]dto.BudgetTransactionResponse, len(yb.BudgetTransactions))
	for i, bt := range yb.BudgetTransactions {
		transactions[i] = dto.BudgetTransactionResponse{
			ID:        bt.ID.String(),
			Name:      bt.Name,
			Amount:    bt.Amount,
			Type:      string(bt.Type),
			CreatedAt: bt.CreatedAt.Unix(),
		}
	}
	return dto.YearlyBudgetResponse{
		ID:                yb.ID.String(),
		UserID:            yb.UserID.String(),
		Year:              yb.Year,
		TotalExpenditures: yb.TotalExpenditures,
		TotalIncome:       yb.TotalIncome,
		TotalSavings:      yb.TotalSavings,
		TotalTransactions: yb.TotalTransactions,
		Transactions:      transactions,
		CreatedAt:         yb.CreatedAt.Unix(),
		UpdatedAt:         yb.UpdatedAt.Unix(),
	}
}
