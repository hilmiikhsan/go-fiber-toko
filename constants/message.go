package constants

import "errors"

var (
	ErrUserNotFound         = errors.New("User not found")
	ErrPasswordNotMatch     = errors.New("Password is not match")
	ErrRecordNotFound       = errors.New("record not found")
	ErrNamaTokoIsRequired   = errors.New("Nama toko is required")
	ErrTokoNotFound         = errors.New("Toko tidak ditemukan")
	ErrProductNotFound      = errors.New("Produk tidak ditemukan")
	ErrCategoryNotFound     = errors.New("Category not found")
	ErrNamaProdukIsRequired = errors.New("Nama Produk is required")
	ErrDeskripsiIsRequired  = errors.New("Deskripsi is required")
)
