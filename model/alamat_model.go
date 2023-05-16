package model

type AlamatModel struct {
	JudulAlamat  string `json:"judul_alamat" validate:"required"`
	NamaPenerima string `json:"nama_penerima" validate:"required"`
	NoTelp       string `json:"no_telp" validate:"required"`
	DetailAlamat string `json:"detail_alamat" validate:"required"`
}

type ParamsModel struct {
	JudulAlamat string `query:"judul_alamat"`
}

type GetAllAlamatModel struct {
	ID           int    `json:"id"`
	JudulAlamat  string `json:"judul_alamat"`
	NamaPenerima string `json:"nama_penerima"`
	NoTelp       string `json:"no_telp"`
	DetailAlamat string `json:"detail_alamat"`
}
