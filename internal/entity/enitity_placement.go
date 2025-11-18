package entity

import "time"

type Placement struct {
	ID          string `gorm:"column:id;primaryKey"` // pakai varchar
	BlockID     string `gorm:"column:block_id"`
	ContainerID string `gorm:"column:container_id"`
	SlotStart   int    `gorm:"column:slot_start"`
	SlotEnd     int    `gorm:"column:slot_end"`
	RowNum      int    `gorm:"column:row_num"`
	Tier        int    `gorm:"column:tier"`
	Status      string `gorm:"type:varchar(50)"`
	CreatedAt   time.Time
}
