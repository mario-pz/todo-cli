CREATE DATABASE IF NOT EXISTS tododb;

USE tododb;

GRANT SELECT, INSERT, UPDATE, DELETE ON tododb.* TO 'user'@'%';

CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    INDEX idx_username (username)
);

CREATE TABLE IF NOT EXISTS tasks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    description VARCHAR(255) NOT NULL,
    completed BOOL NOT NULL DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_user_id (user_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

INSERT INTO users (username, password) VALUES
    ('jdoe', 'abc123'),
    ('asmith', 'qwerty'),
    ('krogers', 'mypassword'),
    ('davidmartinez@yahoo.gr', 'letmein');

INSERT INTO tasks (user_id, description, completed) VALUES 
    (1, 'Buy groceries', false),
    (1, 'Finish homework', false),
    (2, 'Go to gym', true),
    (2, 'Attend meeting', false),
    (4, 'Date with rebecca', false);
