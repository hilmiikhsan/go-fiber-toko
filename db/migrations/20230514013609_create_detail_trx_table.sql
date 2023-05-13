-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS detail_trx
(
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    id_trx INTEGER NOT NULL,
    id_log_produk INTEGER NOT NULL,
    id_toko INTEGER NOT NULL,
    kuantitas INTEGER NOT NULL,
    harga_total INTEGER NOT NULL,
    updated_at DATE NOT NULL,
    created_at DATE NOT NULL,
    FOREIGN KEY (id_trx) REFERENCES trx(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (id_log_produk) REFERENCES log_produk(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (id_toko) REFERENCES toko(id) ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS detail_trx;
-- +goose StatementEnd
