package model

type GetDetailTrxModel struct {
	Product    GetLogProductModel `json:"product"`
	Toko       GetTokoModel       `json:"toko"`
	Kuantitas  int                `json:"kuantitas"`
	HargaTotal int                `json:"harga_total"`
}
