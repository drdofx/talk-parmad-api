package database

import (
	"fmt"

	"github.com/drdofx/talk-parmad/internal/api/lib"
	"github.com/drdofx/talk-parmad/internal/api/models"
	"go.uber.org/fx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Module = fx.Module("database",
	fx.Provide(NewDatabase),
)

// Database is a struct that contains the database connection
type Database struct {
	DB *gorm.DB
}

// NewDatabase creates a new database connection
func NewDatabase(env *lib.Env) (*Database, error) {
	// Connect to the mysql database
	user := env.DBUsername
	password := env.DBPassword
	host := env.DBHost
	port := env.DBPort
	dbname := env.DBName

	fmt.Println(user, password, host, port, dbname)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	database := &Database{db}

	// Migrate the database
	err = database.migrateDB()

	if err != nil {
		return nil, err
	}

	return database, nil

}

// Close closes the database connection
func (d *Database) Close() error {
	db, err := d.DB.DB()

	if err != nil {
		return err
	}

	err = db.Close()

	if err != nil {
		return err
	}

	return nil
}

// migrateDB migrates the database
func (d *Database) migrateDB() error {
	err := d.DB.AutoMigrate(
		models.User{},
		models.Forum{},
		models.UserForum{},
		models.Moderator{},
		models.Thread{},
		models.Reply{},
	)

	if err != nil {
		return err
	}

	return nil
}
