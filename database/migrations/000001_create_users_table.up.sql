CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) NULL,
    firstname VARCHAR(100) NOT NULL,
    lastname VARCHAR(100),
    password VARCHAR(100) NOT NULL,
    age INT,
    address VARCHAR(255),
    role ENUM('customer', 'owner', 'admin', 'player') DEFAULT 'customer',
    email VARCHAR(100) UNIQUE,
    phone VARCHAR(15),
    skill_level ENUM('beginner', 'intermediate', 'advanced') DEFAULT 'beginner',
    position ENUM('goalkeeper', 'defender', 'midfielder', 'forward') DEFAULT 'midfielder',
    profile_picture_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB;
