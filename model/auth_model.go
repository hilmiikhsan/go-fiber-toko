package model

type AuthRegisterModel struct {
	Nama         string `json:"nama"`
	KataSandi    string `json:"kata_sandi"`
	NoTelp       string `json:"no_telp"`
	TanggalLahir string `json:"tanggal_lahir"`
	JenisKelamin string `json:"jenis_kelamin"`
	Tentang      string `json:"tentang"`
	Pekerjaan    string `json:"pekerjaan"`
	Email        string `json:"email"`
	IdProvinsi   string `json:"id_provinsi"`
	IdKota       string `json:"id_kota"`
}

type AuthLoginModel struct {
	NoTelp    string `json:"no_telp"`
	KataSandi string `json:"kata_sandi"`
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
