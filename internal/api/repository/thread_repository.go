package repository

import (
	"github.com/drdofx/talk-parmad/internal/api/database"
	"github.com/drdofx/talk-parmad/internal/api/models"
	"github.com/drdofx/talk-parmad/internal/api/request"
)

type ThreadRepository interface {
	GetThreadByID(id uint) (*models.Thread, error)
	CreateThread(req *request.ReqSaveThread, forumID uint, userID uint) (*models.Thread, error)
	CreateOrUpdateThreadVote(thread *models.Thread, req *request.ReqVoteThread, userID uint) (*models.ThreadVote, error)
	UpdateThread(thread *models.Thread, req *request.ReqEditThread) (*models.Thread, error)
}

type threadRepository struct {
	db *database.Database
}

func NewThreadRepository(db *database.Database) ThreadRepository {
	return &threadRepository{db}
}

func (r *threadRepository) CreateThread(req *request.ReqSaveThread, forumID uint, userID uint) (*models.Thread, error) {
	thread := models.Thread{
		ForumID:   forumID,
		CreatedBy: userID,
		Title:     req.Title,
		Text:      req.Text,
	}

	err := r.db.DB.Create(&thread).Error
	if err != nil {
		return nil, err
	}

	return &thread, nil
}

func (r *threadRepository) GetThreadByID(id uint) (*models.Thread, error) {
	var thread models.Thread
	err := r.db.DB.Where("id = ?", id).First(&thread).Error
	if err != nil {
		return nil, err
	}

	return &thread, nil
}

func (r *threadRepository) CreateOrUpdateThreadVote(thread *models.Thread, req *request.ReqVoteThread, userID uint) (*models.ThreadVote, error) {
	var threadVote models.ThreadVote

	// first, try to update
	update := r.db.DB.Model(&threadVote).Where("thread_id = ?", thread.ID).Where("user_id = ?", userID).Scan(&threadVote).Update("vote", req.Vote)

	if update.RowsAffected == 0 {
		// if no rows affected, create
		threadVote = models.ThreadVote{
			ThreadID: thread.ID,
			UserID:   userID,
			Vote:     req.Vote,
		}

		err := r.db.DB.Create(&threadVote).Error
		if err != nil {
			return nil, err
		}
	}

	return &threadVote, nil
}

func (r *threadRepository) UpdateThread(thread *models.Thread, req *request.ReqEditThread) (*models.Thread, error) {
	err := r.db.DB.Model(&thread).Updates(&req).Error

	if err != nil {
		return nil, err
	}

	return thread, nil
}
