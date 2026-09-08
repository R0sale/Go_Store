-- +goose Up
-- +goose StatementBegin
CREATE TABLE catalog (
    id SERIAL PRIMARY KEY,
    name TEXT,
    price NUMERIC(10, 2),
    image_url TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS catalog;
-- +goose StatementEnd
