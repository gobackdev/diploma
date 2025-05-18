package order

import "diploma/internal/models"

type Repository interface {
	IsOrderExistsForUser(userID uint, orderNumber string) (bool, error)
	IsOrderExistsForOther(userID uint, orderNumber string) (bool, error)
	CreateOrder(userID uint, orderNumber string) error
	GetOrdersByUser(userId uint) ([]models.Order, error)
}
