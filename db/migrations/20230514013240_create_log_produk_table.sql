-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS log_produk
(
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    id_produk INTEGER NOT NULL,
    nama_produk VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL,
    harga_reseller VARCHAR(255) NOT NULL,
    harga_konsumen VARCHAR(255) NOT NULL,
    deskripsi TEXT NOT NULL,
    updated_at DATE NOT NULL,
    created_at DATE NOT NULL,
    id_toko INTEGER NOT NULL,
    id_category INTEGER NOT NULL,
    FOREIGN KEY (id_produk) REFERENCES produk(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (id_toko) REFERENCES toko(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (id_category) REFERENCES category(id) ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS log_produk;
-- +goose StatementEnd
