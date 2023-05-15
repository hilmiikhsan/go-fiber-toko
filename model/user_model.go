package model

type User struct {
	Nama         string   `json:"nama"`
	NoTelp       string   `json:"no_telp"`
	TanggalLahir string   `json:"tanggal_lahir"`
	Tentang      string   `json:"tentang"`
	Pekerjaan    string   `json:"pekerjaan"`
	Email        string   `json:"email"`
	IdProvinsi   Provinsi `json:"id_provinsi"`
	IdKota       Kota     `json:"id_kota"`
}
