package model

type ProductModel struct {
	NamaProduk    string
	CategoryID    int
	HargaReseller int
	HargaKonsumen int
	Stok          int
	Deskripsi     string
}

type ParamsProductModel struct {
	NamaProduk string `query:"nama_produk"`
	Limit      int    `query:"limit"`
	Page       int    `query:"page"`
	CategoryID int    `query:"category_id"`
	TokoID     int    `query:"toko_id"`
	MaxHarga   int    `query:"max_harga"`
	MinHarga   int    `query:"min_harga"`
}

type GetProductModel struct {
	ID            int               `json:"id"`
	NamaProduk    string            `json:"nama_produk"`
	Slug          string            `json:"slug"`
	HargaReseller int               `json:"harga_reseller"`
	HargaKonsumen int               `json:"harga_konsumen"`
	Stok          int               `json:"stok"`
	Deskripsi     string            `json:"deskripsi"`
	Toko          GetTokoModel      `json:"toko"`
	Category      GetCategoryModel  `json:"category"`
	Photos        []FotoProdukModel `json:"photos"`
}
