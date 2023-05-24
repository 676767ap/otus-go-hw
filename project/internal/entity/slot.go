package entity

import (
	"time"

	"gorm.io/gorm"
)

type Slot struct {
	ID        int32     `json:"id"`
	Text      string    `gorm:"not null" json:"text"`
	CreatedAt time.Time `gorm:"not null" json:"-"`
	UpdatedAt time.Time `gorm:"not null" json:"-"`
	Banners   []*Banner `gorm:"many2many:banner_slots;"`
	DeletedAt gorm.DeletedAt
}
