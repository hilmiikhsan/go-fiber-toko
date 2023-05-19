package model

type CategoryModel struct {
	NamaCategory string `json:"nama_category" validate:"required"`
}
