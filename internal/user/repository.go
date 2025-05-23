package user

import "diploma/internal/models"

type Repository interface {
	IsUserExists(login string) (bool, error)
	CreateUser(login, passwordHash string) error
	FindByLogin(login string) (*models.User, error)
}
