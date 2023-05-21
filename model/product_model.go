package model

type CreateProductModel struct {
	NamaProduk    string
	CategoryID    int
	HargaReseller int
	HargaKonsumen int
	Stok          int
	Deskripsi     string
}
