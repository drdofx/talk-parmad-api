package repository

import (
	"github.com/drdofx/talk-parmad/internal/api/database"
	"github.com/drdofx/talk-parmad/internal/api/models"
	"github.com/drdofx/talk-parmad/internal/api/request"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserByNIM(nim *string) (*models.User, error)
	Create(req *request.ReqSaveUser) (*models.User, error)
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

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}

	if err := r.db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetUserByNIM(nim *string) (*models.User, error) {
	user := &models.User{}

	if err := r.db.DB.Where("nim = ?", &nim).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Create(req *request.ReqSaveUser) (*models.User, error) {
	user := &models.User{
		Name:     "Test",
		Email:    req.Email,
		NIM:      &req.NIM,
		Password: req.Password,
	}

	err := r.db.DB.Create(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil

}
