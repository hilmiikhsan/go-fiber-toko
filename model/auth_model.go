package model

type AuthRegisterModel struct {
	Nama         string `json:"nama" validate:"required"`
	KataSandi    string `json:"kata_sandi" validate:"required"`
	NoTelp       string `json:"no_telp" validate:"required"`
	TanggalLahir string `json:"tanggal_lahir" validate:"required"`
	JenisKelamin string `json:"jenis_kelamin" validate:"required"`
	Tentang      string `json:"tentang" validate:"required"`
	Pekerjaan    string `json:"pekerjaan" validate:"required"`
	Email        string `json:"email" validate:"required"`
	IdProvinsi   string `json:"id_provinsi" validate:"required"`
	IdKota       string `json:"id_kota" validate:"required"`
}

type AuthLoginModel struct {
	NoTelp    string `json:"no_telp" validate:"required"`
	KataSandi string `json:"kata_sandi" validate:"required"`
}

type AuthResponseModel struct {
	Nama         string      `json:"nama"`
	NoTelp       string      `json:"no_telp"`
	TanggalLahir string      `json:"tanggal_lahir"`
	Tentang      string      `json:"tentang"`
	Pekerjaan    string      `json:"pekerjaan"`
	Email        string      `json:"email"`
	IdProvinsi   Provinsi    `json:"id_provinsi"`
	IdKota       Kota        `json:"id_kota"`
	Token        interface{} `json:"token"`
}
