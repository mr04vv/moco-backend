
-- +migrate Up
CREATE TABLE IF NOT EXISTS users (id int primary key AUTO_INCREMENT, name varchar(20), profile_url text);

-- +migrate Down
DROP TABLE users;
