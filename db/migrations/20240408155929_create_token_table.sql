-- +goose Up
-- +goose StatementBegin
CREATE TABLE tokens (
  token VARCHAR(255) NOT NULL UNIQUE,
  expires_at INTEGER NOT NULL,
  user_id INTEGER NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tokens;
-- +goose StatementEnd