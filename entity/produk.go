package entity

import "time"

type Produk struct {
	ID            int       `gorm:"primaryKey;column:id;autoIncrement;type:int(11);not null"`
	NamaProduk    string    `gorm:"column:nama_produk;type:varchar(255);not null"`
	Slug          string    `gorm:"column:slug;type:varchar(255);not null"`
	HargaReseller string    `gorm:"column:harga_reseller;type:varchar(255);not null"`
	HargaKonsumen string    `gorm:"column:harga_konsumen;type:varchar(255);not null"`
	Stok          int       `gorm:"column:stok;type:int(11);not null"`
	Deskripsi     string    `gorm:"column:deskripsi;type:text;not null"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:date;not null"`
	CreatedAt     time.Time `gorm:"column:created_at;type:date;not null"`
	IdToko        int       `gorm:"column:id_toko;type:int(11);not null"`
	IdCategory    int       `gorm:"column:id_category;type:int(11);not null"`
	Toko          Toko      `gorm:"ForeignKey:IdToko;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
	Category      Category  `gorm:"ForeignKey:IdCategory;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
}

func (Produk) TableName() string {
	return "produk"
}
