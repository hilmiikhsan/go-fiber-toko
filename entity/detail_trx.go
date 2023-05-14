package entity

import "time"

type DetailTrx struct {
	ID          int       `gorm:"primaryKey;column:id;autoIncrement;type:int(11);not null"`
	IdTrx       int       `gorm:"column:id_trx;type:int(11);not null"`
	IdLogProduk int       `gorm:"column:id_log_produk;type:int(11);not null"`
	IdToko      int       `gorm:"column:id_toko;type:int(11);not null"`
	Kuantitas   int       `gorm:"column:kuantitas;type:int(11);not null"`
	HargaTotal  int       `gorm:"column:harga_total;type:int(11);not null"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:date;not null"`
	CreatedAt   time.Time `gorm:"column:created_at;type:date;not null"`
	Trx         Trx       `gorm:"ForeignKey:IdTrx;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
	LogProduk   LogProduk `gorm:"ForeignKey:IdLogProduk;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
	Toko        Toko      `gorm:"ForeignKey:IdToko;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
}

func (DetailTrx) TableName() string {
	return "detail_trx"
}
