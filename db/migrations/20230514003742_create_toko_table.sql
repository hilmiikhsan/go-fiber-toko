-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS toko
(
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    id_user INTEGER NOT NULL,
    nama_toko VARCHAR(255) NOT NULL,
    url_foto VARCHAR(255) NULL,
    updated_at DATE NOT NULL,
    created_at DATE NOT NULL,
    FOREIGN KEY (id_user) REFERENCES user(id) ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS toko;
-- +goose StatementEnd
