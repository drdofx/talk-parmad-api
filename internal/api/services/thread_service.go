package services

import (
	"fmt"
	"strconv"

	"github.com/drdofx/talk-parmad/internal/api/helper"
	"github.com/drdofx/talk-parmad/internal/api/lib"
	"github.com/drdofx/talk-parmad/internal/api/models"
	"github.com/drdofx/talk-parmad/internal/api/repository"
	"github.com/drdofx/talk-parmad/internal/api/request"
)

type ThreadService interface {
	CreateThread(req *request.ReqSaveThread, user *lib.UserData) (*models.Thread, error)
	VoteThread(req *request.ReqVoteThread, user *lib.UserData) (*models.ThreadVote, error)
	EditThread(req *request.ReqEditThread, user *lib.UserData) (*models.Thread, error)
}

type threadService struct {
	repository      repository.ThreadRepository
	forumRepo       repository.ForumRepository
	transactionRepo repository.TransactionRepository
}

func NewThreadService(
	repository repository.ThreadRepository,
	forumRepo repository.ForumRepository,
	transactionRepo repository.TransactionRepository,
) ThreadService {
	return &threadService{repository, forumRepo, transactionRepo}
}

func (s *threadService) CreateThread(req *request.ReqSaveThread, user *lib.UserData) (*models.Thread, error) {
	tx := s.transactionRepo.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			s.transactionRepo.RollbackTransaction(tx)
		}
	}()

	// Get the forum by id
	forumIdInt, _ := strconv.Atoi(req.ForumID)
	forum, err := s.forumRepo.GetForumById(uint(forumIdInt))
	if err != nil {
		return nil, err
	}

	// Check if user a member of the requested forum
	userForum, _ := s.forumRepo.GetUserForumByID(forum.ID, user.UserID)
	if userForum == nil {
		return nil, fmt.Errorf(helper.UserNotMember)
	}

	// Create the thread for the forum
	createdThread, err := s.repository.CreateThread(req, forum.ID, user.UserID)
	if err != nil {
		return nil, err
	}

	if err := s.transactionRepo.CommitTransaction(tx); err != nil {
		return nil, err
	}

	return createdThread, nil
}

func (s *threadService) VoteThread(req *request.ReqVoteThread, user *lib.UserData) (*models.ThreadVote, error) {
	tx := s.transactionRepo.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			s.transactionRepo.RollbackTransaction(tx)
		}
	}()

	// Get thread by id
	threadIdInt, _ := strconv.Atoi(req.ThreadID)
	thread, err := s.repository.GetThreadByID(uint(threadIdInt))
	if err != nil {
		return nil, err
	}

	// Check if user a member of the requested forum
	userForum, _ := s.forumRepo.GetUserForumByID(thread.ForumID, user.UserID)
	if userForum == nil {
		return nil, fmt.Errorf(helper.UserNotMember)
	}

	// Update the thread data vote
	threadVote, err := s.repository.CreateOrUpdateThreadVote(thread, req, user.UserID)
	if err != nil {
		return nil, err
	}

	if err := s.transactionRepo.CommitTransaction(tx); err != nil {
		return nil, err
	}

	return threadVote, nil
}

func (s *threadService) EditThread(req *request.ReqEditThread, user *lib.UserData) (*models.Thread, error) {
	tx := s.transactionRepo.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			s.transactionRepo.RollbackTransaction(tx)
		}
	}()

	// Get thread by id
	threadIdInt, _ := strconv.Atoi(req.ThreadID)
	thread, err := s.repository.GetThreadByID(uint(threadIdInt))
	if err != nil {
		return nil, err
	}

	// Check if user created the thread
	if thread.CreatedBy != user.UserID {
		return nil, fmt.Errorf(helper.UserNotCreatedThread)
	}

	// Update the thread data
	updatedThread, err := s.repository.UpdateThread(thread, req)
	if err != nil {
		return nil, err
	}

	if err := s.transactionRepo.CommitTransaction(tx); err != nil {
		return nil, err
	}

	return updatedThread, nil
}
