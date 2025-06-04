package database

import (
	"fmt"

	"github.com/cendolin/backend-auth-service/config"
	"github.com/cendolin/backend-auth-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase(config *config.Config) *Database {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", config.Database.User, config.Database.Password, config.Database.DatabaseName, config.Database.Port),
	}))
	if err != nil {
		panic(err)
	}

	return &Database{
		DB: db,
	}
}

// Implementations

func (d *Database) Migrate() {
	if err := d.DB.AutoMigrate(
		&models.User{},
	); err != nil {
		panic(err)
	}
}
