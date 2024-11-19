

-- Tạo bảng `managers`
CREATE TABLE booking_status (
    id CHAR(36) PRIMARY KEY,                             
    name VARCHAR(255),                 
    descriptions VARCHAR(255),                 
    
    -- TS
    deleted_at TIMESTAMP NULL DEFAULT NULL,              
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,      
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, 

    -- Index
    INDEX index_booking_status_created_at (created_at),
    INDEX index_booking_status_updated_at (updated_at),
    INDEX index_booking_status_deleted_at (deleted_at)
);



