package entity

import "time"

type Toko struct {
	ID        int       `gorm:"primaryKey;column:id;autoIncrement;type:int(11);not null"`
	IdUser    int       `gorm:"column:id_user;type:int(11);not null"`
	NamaToko  string    `gorm:"column:nama_toko;type:varchar(255);not null"`
	UrlFoto   string    `gorm:"column:url_foto;type:varchar(255);null"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:date;not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:date;not null"`
	User      User      `gorm:"ForeignKey:IdUser;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
}

func (Toko) TableName() string {
	return "toko"
}
