-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    uuid VARCHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO users (`uuid`, `name`, `email`, `password`) VALUES ("3d6ebd8d-6a04-449d-b232-bd1ca3ead0bf", "ADMIN", "admin@admin.com", "$2a$12$eHHwhkiE6/IEE37g1SLHRe5qcj1bTd26JMqRTdMonrY0hpiYSNJEe");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
