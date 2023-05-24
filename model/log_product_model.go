package model

type GetLogProductModel struct {
	ID            int                `json:"id"`
	NamaProduk    string             `json:"nama_produk"`
	Slug          string             `json:"slug"`
	HargaReseller int                `json:"harga_reseller"`
	HargaKonsumen int                `json:"harga_konsumen"`
	Deskripsi     string             `json:"deskripsi"`
	Toko          GetTokoDetailModel `json:"toko"`
	Category      GetCategoryModel   `json:"category"`
	Photos        []FotoProdukModel  `json:"photos"`
}
