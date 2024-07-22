-- +goose Up
-- +goose StatementBegin
INSERT INTO users (name, email, password, created_at, updated_at) VALUES('rohan', 'singhrohankumar7@gmail.com', '020202', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users WHERE id = 1;
-- +goose StatementEnd
