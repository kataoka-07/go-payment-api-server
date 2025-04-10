SET NAMES utf8mb4;

-- 企業
CREATE TABLE companies (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    ceo_name VARCHAR(255),
    phone VARCHAR(20),
    postal_code VARCHAR(10),
    address TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

-- ユーザー (企業に紐づく)
CREATE TABLE users (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    company_id BIGINT NOT NULL,
    name VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    password TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (company_id) REFERENCES companies(id)
) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

-- 取引先 (企業に紐づく)
CREATE TABLE partners (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    company_id BIGINT NOT NULL,
    name VARCHAR(255),
    ceo_name VARCHAR(255),
    phone VARCHAR(20),
    postal_code VARCHAR(10),
    address TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (company_id) REFERENCES companies(id)
) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

-- 取引先銀行口座 (取引先に紐づく)
CREATE TABLE partner_bank_accounts (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    partner_id BIGINT NOT NULL,
    bank_name VARCHAR(255),
    branch_name VARCHAR(255),
    account_number VARCHAR(50),
    account_holder_name VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (partner_id) REFERENCES partners(id)
) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

-- 請求書データ (企業・取引先に紐づく)
CREATE TABLE invoices (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    company_id BIGINT NOT NULL,
    partner_id BIGINT NOT NULL,
    issue_date DATE NOT NULL,
    payment_amount BIGINT NOT NULL,
    fee BIGINT NOT NULL,
    fee_rate FLOAT NOT NULL,
    tax BIGINT NOT NULL,
    tax_rate FLOAT NOT NULL,
    total_amount BIGINT NOT NULL,
    due_date DATE NOT NULL,
    status ENUM('pending', 'processing', 'paid', 'error') DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (company_id) REFERENCES companies(id),
    FOREIGN KEY (partner_id) REFERENCES partners(id)
) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
