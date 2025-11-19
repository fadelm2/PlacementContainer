package entity

type Yard struct {
	ID     string  `gorm:"primaryKey;column:id"`
	Code   string  `gorm:"unique;not null;column:code"`
	Name   string  `gorm:"column:name"`
	Blocks []Block `gorm:"foreignKey:YardID"`
}

func (l *Yard) TableName() string {
	return "yard"
}
