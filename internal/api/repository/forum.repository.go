package repository

import "github.com/drdofx/talk-parmad/internal/api/database"

type ForumRepository interface {
}

type forumRepository struct {
	db *database.Database
}

func NewForumRepository(db *database.Database) ForumRepository {
	return &forumRepository{db}
}
