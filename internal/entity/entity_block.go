package entity

func (l *Block) Block() string {
	return "Block"
}

type Block struct {
	ID        string `gorm:"column:id;primaryKey"`
	YardID    string `gorm:"column:yard_id"`
	Name      string `gorm:"column:name"`
	TotalSlot int    `gorm:"column:total_slot"`
	TotalRow  int    `gorm:"column:total_row"`
	TotalTier int    `gorm:"column:total_tier"`
}
