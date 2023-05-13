-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS trx
(
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    id_user INTEGER NOT NULL,
    alamat_pengiriman INTEGER NOT NULL,
    harga_total INTEGER NOT NULL,
    kode_invoice VARCHAR(255) NOT NULL,
    method_bayar VARCHAR(255) NOT NULL,
    updated_at DATE NOT NULL,
    created_at DATE NOT NULL,
    FOREIGN KEY (id_user) REFERENCES user(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (alamat_pengiriman) REFERENCES alamat(id) ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop TABLE IF EXISTS trx;
-- +goose StatementEnd
