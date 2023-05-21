package repository

import (
	"github.com/drdofx/talk-parmad/internal/api/database"
	"github.com/drdofx/talk-parmad/internal/api/models"
)

type UserRepository interface {
	Create(user *models.User) (*models.User, error)
	// ReadById(id uint) (*models.User, error)
	// ReadByUsername(username string) (*models.User, error)
	// Update(user *models.User) (*models.User, error)
	// Delete(user *models.User) error
}

type userRepository struct {
	db *database.Database
}

func NewUserRepository(db *database.Database) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *models.User) (*models.User, error) {
	err := r.db.DB.Save(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil

}
