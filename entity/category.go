package entity

import "time"

type Category struct {
	ID           int       `gorm:"primaryKey;column:id;autoIncrement;type:int(11);not null"`
	NamaCategory string    `gorm:"column:name_category;type:varchar(255);not null"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:date;not null"`
	CreatedAt    time.Time `gorm:"column:created_at;type:date;not null"`
}

func (Category) TableName() string {
	return "category"
}
