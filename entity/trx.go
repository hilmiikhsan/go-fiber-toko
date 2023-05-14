package entity

import "time"

type Trx struct {
	ID               int       `gorm:"primaryKey;column:id;autoIncrement;type:int(11);not null"`
	IdUser           int       `gorm:"column:id_user;type:int(11);not null"`
	AlamatPengiriman int       `gorm:"column:alamat_pengiriman;type:int(11);not null"`
	HargaTotal       int       `gorm:"column:harga_total;type:int(11);not null"`
	KodeInvoice      string    `gorm:"column:kode_invoice;type:varchar(255);not null"`
	MethodBayar      string    `gorm:"column:method_bayar;type:varchar(255);not null"`
	UpdatedAt        time.Time `gorm:"column:updated_at;type:date;not null"`
	CreatedAt        time.Time `gorm:"column:created_at;type:date;not null"`
	User             User      `gorm:"ForeignKey:IdUser;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
	Alamat           Alamat    `gorm:"ForeignKey:AlamatPengiriman;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
}

func (Trx) TableName() string {
	return "trx"
}
