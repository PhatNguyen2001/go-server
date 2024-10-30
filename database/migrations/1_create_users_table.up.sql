

-- Tạo bảng `users`
CREATE TABLE users (
    id CHAR(36) PRIMARY KEY,                             -- ID của bản ghi, UUID (có thể thay CHAR(36) bằng loại UUID nếu DB hỗ trợ)
    phone VARCHAR(255) NOT NULL UNIQUE,                  -- Số điện thoại, phải là duy nhất
    email VARCHAR(255) NOT NULL UNIQUE,                  -- Email, phải là duy nhất
    password VARCHAR(180) NOT NULL,                      -- Mật khẩu
    name VARCHAR(255) NOT NULL,                          -- Tên của người dùng
    remember_me_token VARCHAR(255),                      -- Token remember me, cho phép null
    role_id VARCHAR(255),                                -- ID của vai trò (role), liên kết với bảng roles (chỉ định FK sau nếu cần)
    
    -- TS
    deleted_at TIMESTAMP NULL DEFAULT NULL,              
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,      
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, 

    -- Index
    INDEX index_users_on_role_id (role_id),
    INDEX index_users_created_at (created_at),
    INDEX index_users_updated_at (updated_at),
    INDEX index_users_deleted_at (deleted_at)
);
