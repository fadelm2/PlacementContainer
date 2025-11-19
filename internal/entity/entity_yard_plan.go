package entity

type YardPlan struct {
	ID              string  `gorm:"primaryKey;column:id"`
	YardID          string  `gorm:"column:yard_id"`
	BlockID         string  `gorm:"column:block_id"`
	FromSlot        int     `gorm:"column:from_slot"`
	ToSlot          int     `gorm:"column:to_slot"`
	FromRow         int     `gorm:"column:from_row"`
	ToRow           int     `gorm:"column:to_row"`
	ContainerSize   int     `gorm:"column:container_size"` // 20 or 40
	ContainerHeight float64 `gorm:"column:container_height"`
	ContainerType   string  `gorm:"column:container_type"`
}

func (y *YardPlan) YardPlan() string {
	return "Block"
}
