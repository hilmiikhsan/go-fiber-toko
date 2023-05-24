package model

type TrxModel struct {
	MethodBayar string           `json:"method_bayar" validate:"required"`
	AlamatKirim int              `json:"alamat_kirim" validate:"required"`
	DetailTrx   []DetailTrxModel `json:"detail_trx" validate:"required"`
}

type DetailTrxModel struct {
	ProductID int `json:"product_id" validate:"required"`
	Kuantitas int `json:"kuantitas" validate:"required"`
}

type ParamsTrxModel struct {
	Search string `query:"search"`
	Limit  int    `query:"limit"`
	Page   int    `query:"page"`
}

type GetTrxModel struct {
	ID          int                 `json:"id"`
	HargaTotal  int                 `json:"harga_total"`
	KodeInvoice string              `json:"kode_invoice"`
	MethodBayar string              `json:"method_bayar"`
	AlamatKirim GetAlamatModel      `json:"alamat_kirim"`
	DetailTrx   []GetDetailTrxModel `json:"detail_trx"`
}
