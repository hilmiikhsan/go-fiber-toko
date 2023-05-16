package model

type TokoModel struct {
	ID       int    `json:"id"`
	NamaToko string `json:"nama_toko"`
	UrlFoto  string `json:"url_foto"`
	UserID   int    `json:"user_id"`
}
