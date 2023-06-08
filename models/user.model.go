package models

import (
	"time"

	"gorm.io/gorm"
)


/*
User hanya bisa mempunyai 1 role dari id yang diminta dari request body
(One to Many relationship)
*/
type User struct {
	ID        uint             `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string           `gorm:"not null" validate:"required"`
	Username  string           `gorm:"unique;not null" validate:"required"`
	Password  string           `gorm:"not null" validate:"required"`
	CreatedAt time.Time        `json:"-"`
	UpdatedAt time.Time        `json:"-"`
	DeletedAt gorm.DeletedAt   `json:"-"`
}
