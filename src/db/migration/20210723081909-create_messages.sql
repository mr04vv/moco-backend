
-- +migrate Up
CREATE TABLE IF NOT EXISTS messages (id int primary key AUTO_INCREMENT, message text, user_id int,
FOREIGN KEY `messages_ibfk_1` (user_id) references users(id) ON DELETE CASCADE ON UPDATE CASCADE);

-- +migrate Down
DROP TABLE messages;
