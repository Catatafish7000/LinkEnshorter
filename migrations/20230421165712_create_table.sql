-- +goose Up
-- +goose StatementBegin
CREATE TABLE URLs
(
    longurl        text        NOT NULL,
    shorturl      varchar(10) NOT NULL,
    created_at timestamp,
    PRIMARY KEY (shorturl)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS URLs;
-- +goose StatementEnd
