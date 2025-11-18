package entity

type Yard struct {
	ID          string `gorm:"column:id;primaryKey"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
}

func (l *Yard) TableName() string {
	return "yard"
}
