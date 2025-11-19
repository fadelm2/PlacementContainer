package entity

//
//type Container struct {
//	ID            string  `gorm:"column:id;primaryKey"`
//	ContainerNo   string  `gorm:"column:container_no"`
//	Size          int     `gorm:"column:size;"`                           // 20 / 40
//	Type          string  `gorm:"column:type;size"`                       // dry/reefer/dg
//	Status        string  `gorm:"column:status;size:20"`                  // import/export/transit
//	CurrentSlotID *string `gorm:"column:current_slot_id"`                 // FK nullable
//	CurrentSlot   *Slot   `gorm:"foreignKey:CurrentSlotID;references:ID"` // relasi ke Slot
//}
