

-- Tạo bảng `users`
CREATE TABLE services (
    id CHAR(36) PRIMARY KEY,                             
    name VARCHAR(255),                 
    description VARCHAR(255),                 
    avatar_url VARCHAR(180),                      
    status INT,                      
    type INT,             
    
    -- TS
    deleted_at TIMESTAMP NULL DEFAULT NULL,              
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,      
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, 

    -- Index
    INDEX index_services_created_at (created_at),
    INDEX index_services_updated_at (updated_at),
    INDEX index_services_deleted_at (deleted_at)
);



