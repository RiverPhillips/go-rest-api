-- +goose Up
-- +goose StatementBegin
ALTER TABLE todo ADD completed BOOLEAN NOT NULL DEFAULT FALSE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE todo DROP COLUMN completed;
-- +goose StatementEnd
