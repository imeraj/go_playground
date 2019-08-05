CREATE TABLE
IF NOT EXISTS users
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR
(255) NOT NULL,
    password_hash VARCHAR
(255) NOT NULL,
    email VARCHAR
(255) NOT NULL,
    created_at TIMESTAMP default CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    UNIQUE KEY uix_users_email
(email)
)  ENGINE=INNODB;

