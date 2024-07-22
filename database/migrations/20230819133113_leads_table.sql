-- +goose Up
-- +goose StatementBegin
CREATE TABLE leads (
    id INT NOT NULL AUTO_INCREMENT,
    bussiness_name VARCHAR(45) NOT NULL,
    bussiness_gst VARCHAR(45) NOT NULL,
    contact_number INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE leads;
-- +goose StatementEnd
