-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user
(
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    nama VARCHAR(255) NOT NULL,
    kata_sandi VARCHAR(255) NOT NULL,
    notelp VARCHAR(255) NOT NULL UNIQUE,
    tanggal_lahir DATE NOT NULL,
    jenis_kelamin VARCHAR(255) NOT NULL,
    tentang TEXT NOT NULL,
    pekerikan VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    id_provinsi VARCHAR(255) NOT NULL,
    id_kota VARCHAR(255) NOT NULL,
    isAdmin BOOLEAN NOT NULL,
    updated_at DATE NOT NULL,
    created_at DATE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user;
-- +goose StatementEnd
