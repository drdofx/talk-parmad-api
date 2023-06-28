package services

import (
	"fmt"
	"strconv"

	"github.com/drdofx/talk-parmad/internal/api/helper"
	"github.com/drdofx/talk-parmad/internal/api/lib"
	"github.com/drdofx/talk-parmad/internal/api/models"
	"github.com/drdofx/talk-parmad/internal/api/repository"
	"github.com/drdofx/talk-parmad/internal/api/request"
	"github.com/drdofx/talk-parmad/internal/api/response"
)

type ThreadService interface {
	CreateThread(req *request.ReqSaveThread, user *lib.UserData) (*models.Thread, error)
	VoteThread(req *request.ReqVoteThread, user *lib.UserData) (*models.ThreadVote, error)
	EditThread(req *request.ReqEditThread, user *lib.UserData) (*models.Thread, error)
	DetailThread(req *request.ReqDetailThread) (*response.ResDetailThread, error)
	CreateReply(req *request.ReqSaveReply, user *lib.UserData) (*models.Reply, error)
	VoteReply(req *request.ReqVoteReply, user *lib.UserData) (*models.ReplyVote, error)
	EditReply(req *request.ReqEditReply, user *lib.UserData) (*models.Reply, error)
	ListUserThread(user *lib.UserData) ([]models.Thread, error)
	ListUserReply(user *lib.UserData) ([]*response.ResListThreadReply, error)
	GetThreadByID(threadID uint) (*models.Thread, error)
	CheckModeratorForumFromThread(thread *models.Thread, user *lib.UserData) (bool, error)
	DeleteThread(thread *models.Thread) error
	GetThreadAndReplyByReplyID(replyID uint) (*models.Thread, *models.Reply, error)
	DeleteReply(reply *models.Reply) error
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

func (s *threadService) ListUserThread(user *lib.UserData) ([]models.Thread, error) {
	// Get the thread data
	threads, err := s.repository.ListUserThread(user)
	if err != nil {
		return nil, err
	}

	return threads, nil
}

func (s *threadService) ListUserReply(user *lib.UserData) ([]*response.ResListThreadReply, error) {
	// Get the reply data
	replies, err := s.repository.ListUserReply(user)
	if err != nil {
		return nil, err
	}

	return replies, nil
}

func (s *threadService) DetailThread(req *request.ReqDetailThread) (*response.ResDetailThread, error) {
	// Get the thread data, including its reply
	threadIdInt, _ := strconv.Atoi(req.ThreadID)
	thread, err := s.repository.DetailThread(uint(threadIdInt))
	if err != nil {
		return nil, err
	}

	return thread, nil
}

func (s *threadService) CreateReply(req *request.ReqSaveReply, user *lib.UserData) (*models.Reply, error) {
	tx := s.transactionRepo.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			s.transactionRepo.RollbackTransaction(tx)
		}
	}()

	// Get the thread by id
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

	// Create the reply for the thread
	createdReply, err := s.repository.CreateReply(req, thread.ID, user.UserID)
	if err != nil {
		return nil, err
	}

	if err := s.transactionRepo.CommitTransaction(tx); err != nil {
		return nil, err
	}

	return createdReply, nil
}

func (s *threadService) VoteReply(req *request.ReqVoteReply, user *lib.UserData) (*models.ReplyVote, error) {
	tx := s.transactionRepo.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			s.transactionRepo.RollbackTransaction(tx)
		}
	}()

	// Get reply by id
	replyIdInt, _ := strconv.Atoi(req.ReplyID)
	reply, err := s.repository.GetReplyByID(uint(replyIdInt))
	if err != nil {
		return nil, err
	}

	// Get thread by id
	thread, err := s.repository.GetThreadByID(reply.ThreadID)
	if err != nil {
		return nil, err
	}

	// Check if user a member of the requested forum
	userForum, _ := s.forumRepo.GetUserForumByID(thread.ForumID, user.UserID)
	if userForum == nil {
		return nil, fmt.Errorf(helper.UserNotMember)
	}

	// Update the reply data vote
	replyVote, err := s.repository.CreateOrUpdateReplyVote(reply, req, user.UserID)
	if err != nil {
		return nil, err
	}

	if err := s.transactionRepo.CommitTransaction(tx); err != nil {
		return nil, err
	}

	return replyVote, nil
}

func (s *threadService) EditReply(req *request.ReqEditReply, user *lib.UserData) (*models.Reply, error) {
	tx := s.transactionRepo.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			s.transactionRepo.RollbackTransaction(tx)
		}
	}()

	// Get reply by id
	replyIdInt, _ := strconv.Atoi(req.ReplyID)
	reply, err := s.repository.GetReplyByID(uint(replyIdInt))
	if err != nil {
		return nil, err
	}

	// Check if user created the reply
	if reply.CreatedBy != user.UserID {
		return nil, fmt.Errorf(helper.UserNotCreatedReply)
	}

	// Update the reply data
	updatedReply, err := s.repository.UpdateReply(reply, req)
	if err != nil {
		return nil, err
	}

	if err := s.transactionRepo.CommitTransaction(tx); err != nil {
		return nil, err
	}

	return updatedReply, nil
}

func (s *threadService) GetThreadByID(threadID uint) (*models.Thread, error) {
	thread, err := s.repository.GetThreadByID(threadID)
	if err != nil {
		return nil, err
	}

	return thread, nil
}

func (s *threadService) CheckModeratorForumFromThread(thread *models.Thread, user *lib.UserData) (bool, error) {
	// Check if user a moderator of the requested forum
	moderator, _ := s.forumRepo.GetModeratorByID(thread.ForumID, user.UserID)
	if moderator == nil {
		return false, fmt.Errorf(helper.UserNotModerator)
	}

	return true, nil
}

func (s *threadService) DeleteThread(thread *models.Thread) error {
	// Begin transaction
	tx := s.transactionRepo.BeginTransaction()

	// Defer the rollback in case of an error
	defer func() {
		if r := recover(); r != nil {
			s.transactionRepo.RollbackTransaction(tx)
		}
	}()

	// Delete the thread
	err := s.repository.DeleteThread(thread)

	// Commit the transaction
	s.transactionRepo.CommitTransaction(tx)

	return err
}

func (s *threadService) GetThreadAndReplyByReplyID(replyID uint) (*models.Thread, *models.Reply, error) {
	reply, err := s.repository.GetReplyByID(replyID)
	if err != nil {
		return nil, nil, err
	}

	thread, err := s.repository.GetThreadByID(reply.ThreadID)
	if err != nil {
		return nil, nil, err
	}

	return thread, reply, nil
}

func (s *threadService) DeleteReply(reply *models.Reply) error {
	// Begin transaction
	tx := s.transactionRepo.BeginTransaction()

	// Defer the rollback in case of an error
	defer func() {
		if r := recover(); r != nil {
			s.transactionRepo.RollbackTransaction(tx)
		}
	}()

	// Delete the reply
	err := s.repository.DeleteReply(reply)

	// Commit the transaction
	s.transactionRepo.CommitTransaction(tx)

	return err
}
