package services

import (
	"fmt"

	"github.com/drdofx/talk-parmad/internal/api/helper"
	"github.com/drdofx/talk-parmad/internal/api/lib"
	"github.com/drdofx/talk-parmad/internal/api/models"
	"github.com/drdofx/talk-parmad/internal/api/repository"
	"github.com/drdofx/talk-parmad/internal/api/request"
	"github.com/drdofx/talk-parmad/internal/api/response"
)

type ForumService interface {
	CreateForum(req *request.ReqSaveForum, user *lib.UserData) (*models.Forum, error)
	JoinForum(req *request.ReqJoinForum, user *lib.UserData) error
	CheckModeratorForum(req *request.ReqCheckModeratorForum) (bool, error)
	EditForum(req *request.ReqEditForum) (*models.Forum, error)
	DeleteForum(req *request.ReqDeleteForum) error
	ListUserForum(user *lib.UserData) ([]models.Forum, error)
	ListThreadForumHome(user *lib.UserData) (*[]response.ResThreadForumHome, error)
	DiscoverForum(user *lib.UserData) ([]models.Forum, error)
	DetailForum(user *lib.UserData, req *request.ReqDetailForum) (*response.ResDetailForum, error)
	RemoveFromForum(req *request.ReqRemoveFromForum) error
	SearchForum(req *request.ReqSearchForum) (*[]response.ResSearchForum, error)
	// ReadById(id uint) (*models.Forum, error)
	// ExitForum(req *request.ReqExitForum) (*models.Forum, error)
}

type forumService struct {
	repository      repository.ForumRepository
	transactionRepo repository.TransactionRepository
}

func NewForumService(repo repository.ForumRepository, transactionRepo repository.TransactionRepository) ForumService {
	return &forumService{repo, transactionRepo}
}

func (s *forumService) CreateForum(req *request.ReqSaveForum, user *lib.UserData) (*models.Forum, error) {
	// Begin transaction
	tx := s.transactionRepo.BeginTransaction()

	// Defer the rollback in case of an error
	defer func() {
		if r := recover(); r != nil {
			s.transactionRepo.RollbackTransaction(tx)
		}
	}()

	// Check if user is authorized to create a forum
	if user.Role != "User" {
		return nil, fmt.Errorf(helper.RoleNotAuthorized)
	}

	// Check if forum with the same name already exists
	existingForum, _ := s.repository.GetForumByName(req.ForumName)
	if existingForum != nil {
		return nil, fmt.Errorf(helper.ForumExists)
	}

	createdForum, err := s.repository.CreateForum(req, user)
	if err != nil {
		return nil, err
	}

	// Create the moderator (head) for the forum
	_, err = s.repository.CreateModeratorHead(createdForum, user)
	if err != nil {
		return nil, err
	}

	// Create the user-forum relation
	_, err = s.repository.CreateUserForum(createdForum, user)
	if err != nil {
		return nil, err
	}

	// Commit the transaction
	s.transactionRepo.CommitTransaction(tx)

	return createdForum, nil
}

func (s *forumService) JoinForum(req *request.ReqJoinForum, user *lib.UserData) error {
	// Begin transaction
	tx := s.transactionRepo.BeginTransaction()

	// Defer the rollback in case of an error
	defer func() {
		if r := recover(); r != nil {
			s.transactionRepo.RollbackTransaction(tx)
		}
	}()

	// Get the forum by id
	forum, err := s.repository.GetForumById(req.ForumID)
	if err != nil {
		return err
	}

	// Check if user is already a member of the forum
	userForum, _ := s.repository.GetUserForumByID(forum.ID, user.UserID)
	if userForum != nil {
		return fmt.Errorf(helper.UserAlreadyMember)
	}

	// Create the user-forum relation
	_, err = s.repository.CreateUserForum(forum, user)

	// Commit the transaction
	s.transactionRepo.CommitTransaction(tx)

	return err
}

func (s *forumService) CheckModeratorForum(req *request.ReqCheckModeratorForum) (bool, error) {
	// Get the forum by id
	forum, err := s.repository.GetForumById(req.ForumID)
	if err != nil {
		return false, err
	}

	// Check if user is a moderator of the forum
	moderator, _ := s.repository.GetModeratorByID(forum.ID, req.UserID)
	if moderator == nil {
		return false, fmt.Errorf(helper.UserNotModerator)
	}

	return true, nil
}

func (s *forumService) ListUserForum(user *lib.UserData) ([]models.Forum, error) {
	// Get the list of forums
	forums, err := s.repository.ListUserForum(user)
	if err != nil {
		return nil, err
	}

	return forums, nil
}

func (s *forumService) DiscoverForum(user *lib.UserData) ([]models.Forum, error) {
	// Get the list of not joined forums
	forums, err := s.repository.DiscoverForum(user)
	if err != nil {
		return nil, err
	}

	return forums, nil
}

func (s *forumService) DetailForum(user *lib.UserData, req *request.ReqDetailForum) (*response.ResDetailForum, error) {
	// Get the forum detail, including the list of threads
	forum, err := s.repository.DetailForum(user, req.ForumID)
	if err != nil {
		return nil, err
	}

	// check if user is a member of the forum
	userForum, _ := s.repository.GetUserForumByID(req.ForumID, user.UserID)
	if userForum == nil {
		forum.IsMember = false
	} else {
		forum.IsMember = true
	}

	return forum, nil

}

func (s *forumService) ListThreadForumHome(user *lib.UserData) (*[]response.ResThreadForumHome, error) {
	// Get the list of threads
	threads, err := s.repository.ListThreadForumHome(user.UserID)
	if err != nil {
		return nil, err
	}

	return threads, nil
}

func (s *forumService) EditForum(req *request.ReqEditForum) (*models.Forum, error) {
	// Begin transaction
	tx := s.transactionRepo.BeginTransaction()

	// Defer the rollback in case of an error
	defer func() {
		if r := recover(); r != nil {
			s.transactionRepo.RollbackTransaction(tx)
		}
	}()

	// Get the forum by id
	forum, err := s.repository.GetForumById(req.ForumID)
	if err != nil {
		return nil, err
	}

	// Update the forum
	updatedForum, err := s.repository.UpdateForum(forum, req)

	// Commit the transaction
	s.transactionRepo.CommitTransaction(tx)

	return updatedForum, err
}

func (s *forumService) DeleteForum(req *request.ReqDeleteForum) error {
	// Begin transaction
	tx := s.transactionRepo.BeginTransaction()

	// Defer the rollback in case of an error
	defer func() {
		if r := recover(); r != nil {
			s.transactionRepo.RollbackTransaction(tx)
		}
	}()

	// Get the forum by id
	forum, err := s.repository.GetForumById(req.ForumID)
	if err != nil {
		return err
	}

	// Delete the forum
	err = s.repository.DeleteForum(forum)

	// Commit the transaction
	s.transactionRepo.CommitTransaction(tx)

	return err
}

func (s *forumService) RemoveFromForum(req *request.ReqRemoveFromForum) error {
	// Begin transaction
	tx := s.transactionRepo.BeginTransaction()

	// Defer the rollback in case of an error
	defer func() {
		if r := recover(); r != nil {
			s.transactionRepo.RollbackTransaction(tx)
		}
	}()

	// Check if user is indeed a member of the forum
	userForum, _ := s.repository.GetUserForumByID(req.ForumID, req.UserID)
	if userForum == nil {
		return fmt.Errorf(helper.UserNotMember)
	}

	// Delete the user-forum relation
	err := s.repository.RemoveFromForum(userForum)
	if err != nil {
		return err
	}

	// Commit the transaction
	s.transactionRepo.CommitTransaction(tx)

	return nil
}

func (s *forumService) SearchForum(req *request.ReqSearchForum) (*[]response.ResSearchForum, error) {
	// Get the list of forums
	forums, err := s.repository.SearchForum(req)

	if err != nil {
		return nil, err
	}

	return forums, nil
}
