package model

type CategoryModel struct {
	NamaCategory string `json:"nama_category" validate:"required"`
}

type GetCategoryModel struct {
	ID           int    `json:"id"`
	NamaCategory string `json:"nama_category"`
}
