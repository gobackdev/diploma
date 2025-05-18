package postgres

import "diploma/internal/models"

func (s *Postgres) IsOrderExistsForUser(userID uint, orderNumber string) (bool, error) {
	var count int64
	err := s.DB.Model(&models.Order{}).
		Where("user_id = ? AND order_number = ?", userID, orderNumber).
		Count(&count).Error
	return count > 0, err
}

func (s *Postgres) IsOrderExistsForOther(userID uint, orderNumber string) (bool, error) {
	var count int64
	err := s.DB.Model(&models.Order{}).
		Where("user_id <> ? AND order_number = ?", userID, orderNumber).
		Count(&count).Error
	return count > 0, err
}

func (s *Postgres) CreateOrder(userID uint, orderNumber string) error {
	order := &models.Order{
		UserID:      userID,
		OrderNumber: orderNumber,
	}
	return s.DB.Create(order).Error
}

func (s *Postgres) GetOrdersByUser(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := s.DB.
		Where("user_id = ?", userID).
		Order("uploaded_at DESC").
		Find(&orders).Error
	return orders, err
}
