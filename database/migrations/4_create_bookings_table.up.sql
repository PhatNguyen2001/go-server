

-- Tạo bảng `users`
CREATE TABLE bookings (
    id CHAR(36) PRIMARY KEY,                             
    booking_date VARCHAR(255),                 
    booking_hour_start VARCHAR(255),                 
    booking_hour_end VARCHAR(180),    
                      
    customer_id VARCHAR(255),  
    doctor_id VARCHAR(255),   
    service_id VARCHAR(255),
    booking_status_id VARCHAR(255),               
    
    -- TS
    deleted_at TIMESTAMP NULL DEFAULT NULL,              
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,      
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, 

    -- Index
    INDEX index_bookings_on_customer_id (customer_id),
    INDEX index_bookings_on_tasker_id (doctor_id),
    INDEX index_bookings_on_service_id (service_id),
    INDEX index_bookings_on_booking_status_id (booking_status_id),
    INDEX index_bookings_created_at (created_at),
    INDEX index_bookings_updated_at (updated_at),
    INDEX index_bookings_deleted_at (deleted_at)
);



