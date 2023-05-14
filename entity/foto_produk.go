package entity

import "time"

type FotoProduk struct {
	ID        int       `gorm:"primaryKey;column:id;autoIncrement;type:int(11);not null"`
	IdProduk  int       `gorm:"column:id_produk;type:int(11);not null"`
	Url       string    `gorm:"column:url;type:varchar(255);not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:date;not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:date;not null"`
	Produk    Produk    `gorm:"ForeignKey:IdProduk;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
}

func (FotoProduk) TableName() string {
	return "foto_produk"
}
