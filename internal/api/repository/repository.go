package repository

import (
	"github.com/drdofx/talk-parmad/internal/api/database"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Module("repository",
	fx.Provide(
		NewUserRepository,
		NewForumRepository,
		NewThreadRepository,
		NewGormTransactionRepository,
	),
)

type TransactionRepository interface {
	BeginTransaction() *gorm.DB
	CommitTransaction(tx *gorm.DB) error
	RollbackTransaction(tx *gorm.DB) error
}

type gormTransactionRepository struct {
	db *database.Database
}

func NewGormTransactionRepository(db *database.Database) TransactionRepository {
	return &gormTransactionRepository{db}
}

func (r *gormTransactionRepository) BeginTransaction() *gorm.DB {
	tx := r.db.DB.Begin()
	return tx
}

func (r *gormTransactionRepository) CommitTransaction(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (r *gormTransactionRepository) RollbackTransaction(tx *gorm.DB) error {
	return tx.Rollback().Error
}
