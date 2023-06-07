package repository

import (
	"errors"

	"github.com/drdofx/talk-parmad/internal/api/database"
	"github.com/drdofx/talk-parmad/internal/api/lib"
	"github.com/drdofx/talk-parmad/internal/api/models"
	"github.com/drdofx/talk-parmad/internal/api/request"
	"github.com/drdofx/talk-parmad/internal/api/response"
)

type ForumRepository interface {
	GetForumByName(name string) (*models.Forum, error)
	GetForumById(id uint) (*models.Forum, error)
	GetUserForumByID(forumID uint, userID uint) (*models.UserForum, error)
	GetModeratorByID(forumID uint, userID uint) (*models.Moderator, error)
	CreateForum(req *request.ReqSaveForum, user *lib.UserData) (*models.Forum, error)
	CreateModeratorHead(forum *models.Forum, user *lib.UserData) (*models.Moderator, error)
	CreateUserForum(forum *models.Forum, user *lib.UserData) (*models.UserForum, error)
	ListUserForum(user *lib.UserData) ([]models.Forum, error)
	DetailForum(forumID uint) (*response.ResDetailForum, error)
	ListThreadForumHome(userID uint) (*[]response.ResThreadForumHome, error)
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

func (r *forumRepository) ListUserForum(user *lib.UserData) ([]models.Forum, error) {
	var forums []models.Forum
	err := r.db.DB.
		Table("user_forums as uf").
		Select("f.*").
		Joins("inner join forums as f on f.id = uf.forum_id").
		Where("uf.user_id = ?", user.UserID).
		Scan(&forums).Error

	if err != nil {
		return nil, err
	}

	return forums, nil
}

func (r *forumRepository) DetailForum(forumID uint) (*response.ResDetailForum, error) {
	var res response.ResDetailForum

	forumQuery := `
		SELECT *
		FROM forums
		WHERE id = ?
		AND deleted_at IS NULL
	`

	// Execute forum query
	forumRows, err := r.db.DB.Raw(forumQuery, forumID).Rows()
	if err != nil {
		return nil, err
	}
	defer forumRows.Close()

	// Scan forum data
	if !forumRows.Next() {
		return nil, errors.New("forum not found")
	}

	err = r.db.DB.ScanRows(forumRows, &res.ForumData)
	if err != nil {
		return nil, err
	}

	threadQuery := `
		SELECT *
		FROM threads
		WHERE forum_id = ?
		AND deleted_at IS NULL
	`

	// Execute thread query
	threadRows, err := r.db.DB.Raw(threadQuery, forumID).Rows()
	if err != nil {
		return nil, err
	}
	defer threadRows.Close()

	// Scan thread data
	for threadRows.Next() {
		var t models.Thread
		err = r.db.DB.ScanRows(threadRows, &t)
		if err != nil {
			return nil, err
		}

		res.ThreadData = append(res.ThreadData, t)
	}

	// Calculate total threads count
	res.TotalThreads = len(res.ThreadData)

	// Calculate number of members
	membersQuery := `
		SELECT COUNT(uf.user_id) AS number_of_members
		FROM user_forums AS uf
		WHERE uf.forum_id = ?
		AND is_removed = 0
	`

	err = r.db.DB.Raw(membersQuery, forumID).Scan(&res.NumberOfMembers).Error
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *forumRepository) ListThreadForumHome(userID uint) (*[]response.ResThreadForumHome, error) {
	var res []response.ResThreadForumHome

	query := `
		SELECT uf.user_id, f.forum_name, f.forum_image, f.id AS forum_id, t.id AS thread_id, t.title, t.text
		FROM user_forums AS uf
		INNER JOIN forums AS f ON f.id = uf.forum_id
		INNER JOIN threads AS t ON t.forum_id = f.id
		WHERE uf.user_id = ?
		ORDER BY t.created_at DESC
	`

	rows, err := r.db.DB.Raw(query, userID).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// loop through all rows
	for rows.Next() {
		// Scan the row data into variables
		err = r.db.DB.ScanRows(rows, &res)
		if err != nil {
			return nil, err
		}
	}

	return &res, nil
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
