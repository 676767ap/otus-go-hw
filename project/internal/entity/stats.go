package entity

import (
	"time"

	"gorm.io/gorm"
)

type Stat struct {
	ID         int32     `json:"id"`
	Type       string    `gorm:"not null" json:"type"`
	CreatedAt  time.Time `gorm:"not null" json:"-"`
	UpdatedAt  time.Time `gorm:"not null" json:"-"`
	SlotID     int32     `gorm:"not null" json:"slot_id"`
	BannerID   int32     `gorm:"not null" json:"banner_id"`
	SocGruopID int32     `gorm:"not null" json:"soc_group_id"`
	DeletedAt  gorm.DeletedAt
}
