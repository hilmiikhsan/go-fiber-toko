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
