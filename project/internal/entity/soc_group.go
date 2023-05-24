package entity

import (
	"time"

	"gorm.io/gorm"
)

type SocGroup struct {
	ID        int32     `json:"id"`
	Text      string    `gorm:"not null" json:"text"`
	CreatedAt time.Time `gorm:"not null" json:"-"`
	UpdatedAt time.Time `gorm:"not null" json:"-"`
	DeletedAt gorm.DeletedAt
}
