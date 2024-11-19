

-- Tạo bảng `users`
CREATE TABLE customers (
    id CHAR(36) PRIMARY KEY,                             
    gender VARCHAR(255),                 
    bod VARCHAR(255),                 
    avatar_url VARCHAR(180),                      
    level VARCHAR(255),                      
    total_bookings VARCHAR(255),  
    referral_code VARCHAR(255),   
    card_id VARCHAR(255),
    user_id VARCHAR(255),               
    
    -- TS
    deleted_at TIMESTAMP NULL DEFAULT NULL,              
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,      
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, 

    -- Index
    INDEX index_customers_on_user_id (user_id),
    INDEX index_customers_on_card_id (card_id),
    INDEX index_customers_created_at (created_at),
    INDEX index_customers_updated_at (updated_at),
    INDEX index_customers_deleted_at (deleted_at)
);



