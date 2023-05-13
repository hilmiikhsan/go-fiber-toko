-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS foto_produk
(
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    id_produk INTEGER NOT NULL,
    url VARCHAR(255) NOT NULL,
    updated_at DATE NOT NULL,
    created_at DATE NOT NULL,
    FOREIGN KEY (id_produk) REFERENCES produk(id) ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS foto_produk;
-- +goose StatementEnd
