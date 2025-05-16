package order

type Repository interface {
	IsOrderExistsForUser(userID uint, orderNumber string) (bool, error)
	IsOrderExistsForOther(userID uint, orderNumber string) (bool, error)
	CreateOrder(userID uint, orderNumber string) error
}
