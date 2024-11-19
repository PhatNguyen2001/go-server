

-- Tạo bảng `users`
CREATE TABLE booking_details (
    id CHAR(36) PRIMARY KEY,                             
    original_price DECIMAL(14, 2),                 
    promotional_price DECIMAL(14, 2),                 
    payment_price DECIMAL(14, 2),    
                      
    booking_id VARCHAR(255),            
    
    -- TS
    deleted_at TIMESTAMP NULL DEFAULT NULL,              
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,      
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, 

    -- Index
    INDEX index_booking_details_on_booking_id (booking_id),
    INDEX index_booking_details_created_at (created_at),
    INDEX index_booking_details_updated_at (updated_at),
    INDEX index_booking_details_deleted_at (deleted_at)
);



