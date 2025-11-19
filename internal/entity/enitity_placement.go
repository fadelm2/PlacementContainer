package entity

import (
	"gorm.io/gorm"
)

type Placement struct {
	gorm.Model
	ID              string  `gorm:"primaryKey;column:id"`
	YardID          string  `gorm:"column:yard_id;index"`
	BlockID         string  `gorm:"column:block_id;index"`
	ContainerNumber string  `gorm:"column:container_number;uniqueIndex"`
	Slot            int     `gorm:"column:slot"` // starting slot
	Row             int     `gorm:"column:row"`
	Tier            int     `gorm:"column:tier"`
	Width           int     `gorm:"column:width"` // NEW: 1 for 20ft, 2 for 40ft
	Size            int     `gorm:"column:size"`  // 20/40
	Type            string  `gorm:"column:type"`
	Height          float64 `gorm:"column:height"`
	CreatedAt       int64   `gorm:"autoCreateTime:milli"`
	UpdatedAt       int64   `gorm:"autoUpdateTime:milli"`
}
