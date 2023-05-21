package services

import (
	"github.com/drdofx/talk-parmad/internal/api/models"
	"github.com/drdofx/talk-parmad/internal/api/repository"
)

type UserService interface {
	CreateUser(user *models.User) error
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{repository}
}

func (s *userService) CreateUser(user *models.User) error {
	_, err := s.repository.Create(user)

	if err != nil {
		return err
	}

	return nil
}
