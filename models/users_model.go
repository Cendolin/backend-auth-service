package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID              string         `gorm:"primaryKey"`
	Email           string         `json:"email" gorm:"column:email"`
	PasswordHash    string         `json:"password_hash" gorm:"column:password_hash"`
	Username        string         `json:"username" gorm:"column:username"`
	Biography       string         `json:"biography" gorm:"column:biography"`
	SuspendedReason sql.NullString `json:"suspended_reason" gorm:"column:suspended_reason"`

	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()

	return
}
