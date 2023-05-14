package entity

import "time"

type Alamat struct {
	ID           int       `gorm:"primaryKey;column:id;autoIncrement;type:int(11);not null"`
	IdUser       int       `gorm:"column:id_user;type:int(11);not null"`
	JudulAlamat  string    `gorm:"column:judul_alamat;type:varchar(255);not null"`
	NamaPenerima string    `gorm:"column:nama_penerima;type:varchar(255);not null"`
	NoTelp       string    `gorm:"column:no_telp;type:varchar(255);not null"`
	DetailAlamat string    `gorm:"column:detail_alamat;type:varchar(255);not null"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:date;not null"`
	CreatedAt    time.Time `gorm:"column:created_at;type:date;not null"`
	User         User      `gorm:"ForeignKey:IdUser;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
}

func (Alamat) TableName() string {
	return "alamat"
}
