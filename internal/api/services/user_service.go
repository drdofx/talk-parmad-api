package services

import (
	"fmt"

	"github.com/drdofx/talk-parmad/internal/api/helper"
	"github.com/drdofx/talk-parmad/internal/api/lib"
	"github.com/drdofx/talk-parmad/internal/api/models"
	"github.com/drdofx/talk-parmad/internal/api/repository"
	"github.com/drdofx/talk-parmad/internal/api/request"
)

type UserService interface {
	CreateUser(req *request.ReqSaveUser) (*models.User, error)
	LoginUser(req *request.ReqLoginUser) (interface{}, error)
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

	err := lib.HashPassword(&req.Password)
	if err != nil {
		return nil, err
	}

	user, err = s.repository.Create(req)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) LoginUser(req *request.ReqLoginUser) (interface{}, error) {
	var user *models.User
	var err error

	// Check if req.User is number or email
	if helper.IsNumeric(req.User) {
		user, err = s.repository.GetUserByNIM(&req.User)
	} else {
		user, err = s.repository.GetUserByEmail(req.User)
	}

	if err != nil {
		return nil, err
	}

	if !lib.ComparePassword(user.Password, req.Password) {
		return nil, fmt.Errorf(helper.FailedLogin)
	}

	token := lib.GenerateJWT(user)
	if token == "" {
		return nil, fmt.Errorf(helper.FailedGenerateToken)
	}

	return token, nil
}
