ALTER TABLE users
ADD COLUMN remember_hash 
VARCHAR
(255) NOT NULL AFTER password_hash;