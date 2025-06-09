package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID              string     `gorm:"primaryKey" json:"id"`
	Email           string     `json:"email" gorm:"uniqueIndex:email_idx,column:email"`
	PasswordHash    string     `json:"password_hash" gorm:"column:password_hash"`
	Username        string     `json:"username" gorm:"uniqueIndex:username_idx,column:username"`
	SuspendedReason NullString `json:"suspended_reason" gorm:"column:suspended_reason"`
	Verified        bool       `json:"verified" gorm:"column:verified"`
	Country         string     `json:"country" gorm:"column:country"`

	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()

	return
}
