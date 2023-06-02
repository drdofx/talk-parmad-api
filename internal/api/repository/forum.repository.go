package repository

import (
	"github.com/drdofx/talk-parmad/internal/api/database"
	"github.com/drdofx/talk-parmad/internal/api/lib"
	"github.com/drdofx/talk-parmad/internal/api/models"
	"github.com/drdofx/talk-parmad/internal/api/request"
)

type ForumRepository interface {
	GetForumByName(name string) (*models.Forum, error)
	GetForumById(id uint) (*models.Forum, error)
	GetUserForumByID(forumID uint, userID uint) (*models.UserForum, error)
	GetModeratorByID(forumID uint, userID uint) (*models.Moderator, error)
	CreateForum(req *request.ReqSaveForum, user *lib.UserData) (*models.Forum, error)
	CreateModeratorHead(forum *models.Forum, user *lib.UserData) (*models.Moderator, error)
	CreateUserForum(forum *models.Forum, user *lib.UserData) (*models.UserForum, error)
	UpdateForum(forum *models.Forum, req *request.ReqEditForum) (*models.Forum, error)
	DeleteForum(forum *models.Forum) error
}

type forumRepository struct {
	db *database.Database
}

func NewForumRepository(db *database.Database) ForumRepository {
	return &forumRepository{db}
}

func (r *forumRepository) GetForumByName(name string) (*models.Forum, error) {
	var forum models.Forum
	err := r.db.DB.Where("forum_name = ?", name).First(&forum).Error
	if err != nil {
		return nil, err
	}

	return &forum, nil
}

func (r *forumRepository) GetForumById(id uint) (*models.Forum, error) {
	var forum models.Forum
	err := r.db.DB.Where("id = ?", id).First(&forum).Error
	if err != nil {
		return nil, err
	}

	return &forum, nil
}

func (r *forumRepository) GetUserForumByID(forumID uint, userID uint) (*models.UserForum, error) {
	var userForum models.UserForum
	err := r.db.DB.Where("user_id = ?", userID).Where("forum_id = ?", forumID).First(&userForum).Error
	if err != nil {
		return nil, err
	}

	return &userForum, nil
}

func (r *forumRepository) GetModeratorByID(forumID uint, userID uint) (*models.Moderator, error) {
	var moderator models.Moderator
	err := r.db.DB.Where("user_id = ?", userID).Where("forum_id = ?", forumID).First(&moderator).Error
	if err != nil {
		return nil, err
	}

	return &moderator, nil
}

func (r *forumRepository) CreateForum(req *request.ReqSaveForum, user *lib.UserData) (*models.Forum, error) {
	forum := &models.Forum{
		ForumName:        req.ForumName,
		IntroductionText: req.IntroductionText,
		Category:         &req.Category,
	}

	err := r.db.DB.Create(&forum).Error

	if err != nil {
		return nil, err
	}

	return forum, nil
}

func (r *forumRepository) CreateModeratorHead(forum *models.Forum, user *lib.UserData) (*models.Moderator, error) {
	moderator := &models.Moderator{
		ForumID:  forum.ID,
		UserID:   user.UserID,
		Rank:     "Head",
		Nickname: &user.Name,
	}

	err := r.db.DB.Create(&moderator).Error

	if err != nil {
		return nil, err
	}

	return moderator, nil
}

func (r *forumRepository) CreateUserForum(forum *models.Forum, user *lib.UserData) (*models.UserForum, error) {
	userForum := &models.UserForum{
		ForumID: forum.ID,
		UserID:  user.UserID,
	}

	err := r.db.DB.Create(&userForum).Error

	if err != nil {
		return nil, err
	}

	return userForum, nil
}

func (r *forumRepository) UpdateForum(forum *models.Forum, req *request.ReqEditForum) (*models.Forum, error) {
	err := r.db.DB.Model(&forum).Updates(&req).Error

	if err != nil {
		return nil, err
	}

	return forum, nil
}

func (r *forumRepository) DeleteForum(forum *models.Forum) error {
	err := r.db.DB.Delete(&forum).Error

	if err != nil {
		return err
	}

	return nil
}
