package entity

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	ReviewDate time.Time
	Rating uint
	Comment string
	Picture string `gorm:"type:longtext"`

	UserID uint
	User *User `gorm:"foreignKey: UserID"`
}