package repositories

import (
	"github.com/esuEdu/reurb-backend/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {

	return &userRepository{
		db: db,
	}
}

func (repo *userRepository) Create(user *models.User) (*models.User, error) {
	if err := repo.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
