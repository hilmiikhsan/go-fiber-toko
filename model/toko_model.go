package model

type TokoModel struct {
	ID       int    `json:"id"`
	NamaToko string `json:"nama_toko"`
	UrlFoto  string `json:"url_foto"`
	UserID   int    `json:"user_id"`
}

type GetTokoModel struct {
	ID       int    `json:"id"`
	NamaToko string `json:"nama_toko"`
	UrlFoto  string `json:"url_foto"`
}

type ParamsTokoModel struct {
	Limit int    `query:"limit"`
	Page  int    `query:"page"`
	Nama  string `query:"nama"`
}

type GetTokoDetailModel struct {
	NamaToko string `json:"nama_toko"`
	UrlFoto  string `json:"url_foto"`
}
