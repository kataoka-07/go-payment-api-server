SET NAMES utf8mb4;

-- 企業
INSERT INTO companies (name, ceo_name, phone, postal_code, address)
VALUES ('テスト株式会社', '代表 太郎', '03-1234-5678', '100-0001', '東京都千代田区1-1-1');

-- ユーザー (企業に紐づく)
INSERT INTO users (company_id, name, email, password)
VALUES (1, 'テストユーザー', 'user@example.com', 'hashed-password');

-- 取引先 (企業に紐づく)
INSERT INTO partners (company_id, name, ceo_name, phone, postal_code, address)
VALUES (1, 'テスト取引先', '取引先 次郎', '03-9876-5432', '100-0002', '東京都港区2-2-2');
