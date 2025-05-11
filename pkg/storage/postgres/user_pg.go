package postgres

import "diploma/internal/models"

func (s *Postgres) IsUserExists(login string) (bool, error) {
	var count int64
	err := s.DB.Model(&models.User{}).Where("login = ?", login).Count(&count).Error
	return count > 0, err
}

func (s *Postgres) CreateUser(login, passwordHash string) error {
	user := models.User{
		Login:        login,
		PasswordHash: passwordHash,
	}
	return s.DB.Create(&user).Error
}

func (s *Postgres) FindByLogin(login string) (*models.User, error) {
	var user models.User
	if err := s.DB.Where("login = ?", login).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
