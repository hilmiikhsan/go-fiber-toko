-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS category
(
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    nama_category VARCHAR(255) NOT NULL,
    updated_at DATE NOT NULL,
    created_at DATE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS category;
-- +goose StatementEnd
