package repository

import (
	"fmt"

	"github.com/drdofx/talk-parmad/internal/api/database"
	"github.com/drdofx/talk-parmad/internal/api/lib"
	"github.com/drdofx/talk-parmad/internal/api/models"
	"github.com/drdofx/talk-parmad/internal/api/request"
	"github.com/drdofx/talk-parmad/internal/api/response"
)

type ThreadRepository interface {
	GetThreadByID(id uint) (*models.Thread, error)
	CreateThread(req *request.ReqSaveThread, forumID uint, userID uint) (*models.Thread, error)
	CreateOrUpdateThreadVote(thread *models.Thread, req *request.ReqVoteThread, userID uint) (*models.ThreadVote, error)
	UpdateThread(thread *models.Thread, req *request.ReqEditThread) (*models.Thread, error)
	DetailThread(threadID uint) (*response.ResDetailThread, error)
	ListUserThread(user *lib.UserData) ([]*response.ResListThread, error)
	ListUserReply(user *lib.UserData) ([]*response.ResListThreadReply, error)
	GetReplyByID(id uint) (*models.Reply, error)
	CreateReply(req *request.ReqSaveReply, threadID uint, userID uint) (*models.Reply, error)
	CreateOrUpdateReplyVote(reply *models.Reply, req *request.ReqVoteReply, userID uint) (*models.ReplyVote, error)
	UpdateReply(reply *models.Reply, req *request.ReqEditReply) (*models.Reply, error)
	DeleteThread(thread *models.Thread) error
	DeleteReply(reply *models.Reply) error
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

func (r *threadRepository) DetailThread(threadID uint) (*response.ResDetailThread, error) {
	var res response.ResDetailThread

	threadQuery := `
		SELECT t.id as id, t.title, t.text, t.created_at, u.name as created_by, SUM(CASE WHEN tv.vote = true THEN 1 ELSE 0 END) as total_upvotes, SUM(CASE WHEN tv.vote = false THEN 1 ELSE 0 END) as total_downvotes
		FROM threads t
		LEFT JOIN users u ON u.id = t.created_by
		LEFT JOIN thread_votes tv ON tv.thread_id = t.id
		WHERE t.id = ? 
		AND t.deleted_at IS NULL
		GROUP BY t.id, u.name
	`

	// Execute the thread query
	threadRows, err := r.db.DB.Raw(threadQuery, threadID).Rows()
	if err != nil {
		return nil, err
	}
	defer threadRows.Close()

	// Scan thread data into threadField variable
	if !threadRows.Next() {
		return nil, fmt.Errorf("thread not found")
	}

	var threadField response.ResThreadField
	err = threadRows.Scan(&threadField.ID, &threadField.Title, &threadField.Text, &threadField.CreatedAt, &res.CreatedBy, &res.TotalUpvotes, &res.TotalDownvotes)
	if err != nil {
		return nil, err
	}

	// Assign threadField to res.ThreadData
	res.ThreadData = threadField

	// Retrieve replies for the thread
	repliesQuery := `
		SELECT r.id as id, r.text, r.created_at, u2.name as created_by, SUM(CASE WHEN rv.vote = true THEN 1 ELSE 0 END) as total_upvotes, SUM(CASE WHEN rv.vote = false THEN 1 ELSE 0 END) as total_downvotes
		FROM replies r
		LEFT JOIN users u2 ON u2.id = r.created_by
		LEFT JOIN reply_votes rv ON rv.reply_id = r.id
		WHERE r.thread_id = ?
		AND r.deleted_at IS NULL
		GROUP BY r.id, u2.name
	`

	// Execute the replies query
	repliesRows, err := r.db.DB.Raw(repliesQuery, threadID).Rows()
	if err != nil {
		return nil, err
	}
	defer repliesRows.Close()

	// Iterate over the replies and append them to the ResDetailThread struct
	for repliesRows.Next() {
		var reply response.ResReplyField
		err := repliesRows.Scan(&reply.ID, &reply.Text, &reply.CreatedAt, &reply.CreatedBy, &reply.TotalUpvotes, &reply.TotalDownvotes)
		if err != nil {
			return nil, err
		}
		res.ReplyData = append(res.ReplyData, reply)
	}

	res.TotalReplies = len(res.ReplyData)
	return &res, nil
}

func (r *threadRepository) ListUserThread(user *lib.UserData) ([]*response.ResListThread, error) {
	var res []*response.ResListThread

	err := r.db.DB.
		Table("threads t").
		Select("t.*, f.forum_name, f.forum_image").
		Joins("LEFT JOIN forums f ON f.id = t.forum_id").
		Where("t.created_by = ?", user.UserID).
		Where("t.deleted_at IS NULL").
		Order("t.created_at DESC").
		Scan(&res).Error

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *threadRepository) ListUserReply(user *lib.UserData) ([]*response.ResListThreadReply, error) {
	var res []*response.ResListThreadReply

	err := r.db.DB.
		Table("replies r").
		Select("t.*, r.*, f.forum_name, f.forum_image").
		Joins("LEFT JOIN threads t ON t.id = r.thread_id").
		Joins("LEFT JOIN forums f ON f.id = t.forum_id").
		Where("r.created_by = ?", user.UserID).
		Where("r.deleted_at IS NULL").
		Order("r.created_at DESC").
		Scan(&res).Error

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *threadRepository) GetReplyByID(id uint) (*models.Reply, error) {
	var reply models.Reply

	err := r.db.DB.Where("id = ?", id).First(&reply).Error
	if err != nil {
		return nil, err
	}

	return &reply, nil
}

func (r *threadRepository) CreateReply(req *request.ReqSaveReply, threadID uint, userID uint) (*models.Reply, error) {
	reply := models.Reply{
		ThreadID:  threadID,
		CreatedBy: userID,
		Text:      req.Text,
	}

	err := r.db.DB.Create(&reply).Error
	if err != nil {
		return nil, err
	}

	return &reply, nil
}

func (r *threadRepository) CreateOrUpdateReplyVote(reply *models.Reply, req *request.ReqVoteReply, userID uint) (*models.ReplyVote, error) {
	var replyVote models.ReplyVote

	// first, try to update
	update := r.db.DB.Model(&replyVote).Where("reply_id = ?", reply.ID).Where("user_id = ?", userID).Scan(&replyVote).Update("vote", req.Vote)

	if update.RowsAffected == 0 {
		// if no rows affected, create
		replyVote = models.ReplyVote{
			ReplyID: reply.ID,
			UserID:  userID,
			Vote:    req.Vote,
		}

		err := r.db.DB.Create(&replyVote).Error
		if err != nil {
			return nil, err
		}
	}

	return &replyVote, nil
}

func (r *threadRepository) UpdateReply(reply *models.Reply, req *request.ReqEditReply) (*models.Reply, error) {
	err := r.db.DB.Model(&reply).Updates(&req).Error

	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (r *threadRepository) DeleteThread(thread *models.Thread) error {
	err := r.db.DB.Delete(&thread).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *threadRepository) DeleteReply(reply *models.Reply) error {
	err := r.db.DB.Delete(&reply).Error

	if err != nil {
		return err
	}

	return nil
}
