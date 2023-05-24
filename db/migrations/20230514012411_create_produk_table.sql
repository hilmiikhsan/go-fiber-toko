-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS produk
(
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    nama_produk VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL,
    harga_reseller INTEGER NOT NULL,
    harga_konsumen INTEGER NOT NULL,
    stok INTEGER NOT NULL,
    deskripsi TEXT NOT NULL,
    updated_at DATE NOT NULL,
    created_at DATE NOT NULL,
    id_toko INTEGER NOT NULL,
    id_category INTEGER NOT NULL,
    FOREIGN KEY (id_toko) REFERENCES toko(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (id_category) REFERENCES category(id) ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS produk;
-- +goose StatementEnd
