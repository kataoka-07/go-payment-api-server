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

-- 請求書データ (企業・取引先に紐づく)
INSERT INTO invoices (
  id, company_id, partner_id, issue_date, due_date, payment_amount, fee,
  fee_rate, tax, tax_rate, total_amount, status
)
VALUES
  (1, 1, 1, '2025-04-10', '2025-04-30', 10000, 400, 0.04, 40, 0.10, 10440, 'pending'),
  (2, 1, 1, '2025-04-11', '2025-04-30', 30000, 1200, 0.04, 120, 0.10, 31320, 'pending'),
  (3, 1, 1, '2025-04-12', '2025-04-30', 10500, 420, 0.04, 42, 0.10, 10962, 'pending'),
  (4, 1, 1, '2025-04-13', '2025-04-30', 2200, 88, 0.04, 8, 0.10, 2296, 'pending'),
  (5, 1, 1, '2025-04-13', '2025-04-29', 2200, 88, 0.04, 8, 0.10, 2296, 'pending'),
  (6, 1, 1, '2025-04-13', '2025-04-28', 4200, 168, 0.04, 16, 0.10, 4384, 'pending'),
  (7, 1, 1, '2025-04-15', '2025-04-25', 8000, 320, 0.04, 32, 0.10, 8352, 'pending');
