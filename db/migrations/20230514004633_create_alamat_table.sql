-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS alamat
(
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    id_user INTEGER NOT NULL,
    judul_alamat VARCHAR(255) NOT NULL,
    nama_penerima VARCHAR(255) NOT NULL,
    no_telp VARCHAR(255) NOT NULL,
    detail_alamat VARCHAR(255) NOT NULL,
    updated_at DATE NOT NULL,
    created_at DATE NOT NULL,
    FOREIGN KEY (id_user) REFERENCES user(id) ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS alamat;
-- +goose StatementEnd
