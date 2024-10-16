-- +goose Up
ALTER TABLE feeds
ADD last_fetch_at TIMESTAMP;

-- +goose Down
DROP TABLE feed_follows;