package order

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
	"time"
)

type Handler struct {
	repo Repository
}

func NewHandler(repo Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) CreateOrder(c *gin.Context) {
	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userID, ok := userIDValue.(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	orderNumber := strings.TrimSpace(string(body))

	if !isAllDigits(orderNumber) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "order number must be digits"})
		return
	}

	if !IsValidLuhn(orderNumber) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid order number"})
		return
	}

	existsForUser, err := h.repo.IsOrderExistsForUser(userID, orderNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}
	if existsForUser {
		c.JSON(http.StatusOK, gin.H{"message": "order number already uploaded"})
		return
	}

	existsForOther, err := h.repo.IsOrderExistsForOther(userID, orderNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}
	if existsForOther {
		c.JSON(http.StatusConflict, gin.H{"error": "order number already uploaded"})
		return
	}

	if err := h.repo.CreateOrder(userID, orderNumber); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "order number accepted"})

}

func (h *Handler) GetOrders(c *gin.Context) {
	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID, ok := userIDValue.(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	orders, err := h.repo.GetOrdersByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}
	if len(orders) == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	result := make([]OrderResponse, 0, len(orders))
	for _, o := range orders {
		resp := OrderResponse{
			Number:     o.OrderNumber,
			Status:     o.Status,
			UploadedAt: o.UploadedAt.Format(time.RFC3339),
		}
		if o.Status == "PROCESSED" && o.Accrual != nil {
			resp.Accrual = o.Accrual
		}
		result = append(result, resp)
	}
	c.JSON(http.StatusOK, result)
}

// isAllDigits проверяем что номер заказа состоит только из цифр
func isAllDigits(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// IsValidLuhn проверяем алгоритмом luhn
func IsValidLuhn(number string) bool {
	var sum int
	alt := false
	for i := len(number) - 1; i >= 0; i-- {
		n := int(number[i] - '0')
		if alt {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
		alt = !alt
	}
	return sum%10 == 0
}
