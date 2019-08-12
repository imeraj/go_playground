CREATE TABLE
IF NOT EXISTS galleries
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    title VARCHAR
(255) NOT NULL,
    created_at TIMESTAMP default CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    
    FOREIGN KEY (user_id) REFERENCES users (id),
    INDEX uix_galleries_user_id (user_id)
)  ENGINE=INNODB;

