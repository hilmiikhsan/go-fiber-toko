package entity

import "time"

type User struct {
	ID           int       `gorm:"primaryKey;column:id;autoIncrement;type:int(11);not null"`
	Name         string    `gorm:"column:name;type:varchar(255);not null"`
	KataSandi    string    `gorm:"column:kata_sandi;type:varchar(255);not null"`
	NoTelp       string    `gorm:"column:notelp;type:varchar(255);not null;unique"`
	TanggalLahir time.Time `gorm:"column:tanggal_lahir;type:date;not null"`
	JenisKelamin string    `gorm:"column:jenis_kelamin;type:varchar(255);not null"`
	Tentang      string    `gorm:"column:tentang;type:text;not null"`
	Pekerjaan    string    `gorm:"column:pekerjaan;type:varchar(255);not null"`
	Email        string    `gorm:"column:email;type:varchar(255);not null"`
	IdProvinsi   string    `gorm:"column:id_provinsi;type:varchar(255);not null"`
	IdKota       string    `gorm:"column:id_kota;type:varchar(255);not null"`
	IsAdmin      bool      `gorm:"column:isAdmin;default:false;not null"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:date;not null"`
	CreatedAt    time.Time `gorm:"column:created_at;type:date;not null"`
	// Toko         []Toko    `gorm:"ForeignKey:IdUser;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
}

func (User) TableName() string {
	return "user"
}
