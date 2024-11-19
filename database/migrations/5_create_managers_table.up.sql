

-- Tạo bảng `managers`
CREATE TABLE managers (
    id CHAR(36) PRIMARY KEY,                             
    gender VARCHAR(255),                 
    bod VARCHAR(255),                 
    avatar_url VARCHAR(180),                                   
    
    user_id VARCHAR(255),               
    -- TS
    deleted_at TIMESTAMP NULL DEFAULT NULL,              
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,      
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, 

    -- Index
    INDEX index_managers_on_user_id (user_id),
    INDEX index_managers_created_at (created_at),
    INDEX index_managers_updated_at (updated_at),
    INDEX index_managers_deleted_at (deleted_at)
);



