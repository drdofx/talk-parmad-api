package services

import (
	"github.com/drdofx/talk-parmad/internal/api/models"
	"github.com/drdofx/talk-parmad/internal/api/repository"
	"github.com/drdofx/talk-parmad/internal/api/request"
)

type ForumService interface {
	Create(req *request.ReqSaveForum) (*models.Forum, error)
	// ReadById(id uint) (*models.Forum, error)
	// JoinForum(req *request.ReqJoinForum) (*models.Forum, error)
	// ExitForum(req *request.ReqExitForum) (*models.Forum, error)
	// EditForum(req *request.ReqEditForum) (*models.Forum, error)
	// DeleteForum(req *request.ReqDeleteForum) (*models.Forum, error)
}

type forumService struct {
	repository repository.ForumRepository
}

func NewForumService(repo repository.ForumRepository) ForumService {
	return &forumService{repo}
}

func (s *forumService) Create(req *request.ReqSaveForum) (*models.Forum, error) {
	// Check if user is mahaasiswa

	// Check if forum exists with same name

	// Create forum and assgin user as moderator (head)

	// finish
	return nil, nil
}
