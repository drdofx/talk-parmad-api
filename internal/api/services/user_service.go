package services

import (
	"fmt"

	"github.com/drdofx/talk-parmad/internal/api/helper"
	"github.com/drdofx/talk-parmad/internal/api/models"
	"github.com/drdofx/talk-parmad/internal/api/repository"
	"github.com/drdofx/talk-parmad/internal/api/request"
)

type UserService interface {
	CreateUser(req *request.ReqSaveUser) (*models.User, error)
	LoginUser(req *request.ReqLoginUser) (*models.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{repository}
}

func (s *userService) CreateUser(req *request.ReqSaveUser) (*models.User, error) {
	user, _ := s.repository.GetUserByEmail(req.Email)
	if user != nil {
		return nil, fmt.Errorf(helper.UserExists)
	}

	user, err := s.repository.Create(req)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) LoginUser(req *request.ReqLoginUser) (*models.User, error) {
	user, err := s.repository.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if user.Password != req.Password {
		return nil, fmt.Errorf(helper.FailedLogin)
	}

	return user, nil
}
