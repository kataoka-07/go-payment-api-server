.PHONY: up restart stop clean fmt run model-gen test-db-init test-schema-init test-seed e2e-test

include .env
export

# Docker DB
up:
	docker compose up --build -d

restart:
	docker compose down -v && docker compose up

stop:
	docker compose stop

clean:
	docker compose down

# Server
run:
	go run ./cmd/payment-server/main.go

# Test DB (for e2e test)
test-db-init:
	MYSQL_PWD=$(DB_PASSWORD) mysql --protocol=TCP -u $(DB_USER) -h $(DB_HOST) -P $(DB_PORT) -e \
	'CREATE DATABASE payment_test DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;'
	
test-schema-init:
	MYSQL_PWD=$(DB_PASSWORD) mysql --protocol=TCP -u $(DB_USER) -h $(DB_HOST) -P $(DB_PORT) $(DB_NAME)_test < ./migrations/schema.sql

test-seed:
	MYSQL_PWD=$(DB_PASSWORD) mysql --protocol=TCP -u $(DB_USER) -h $(DB_HOST) -P $(DB_PORT) $(DB_NAME)_test < ./migrations/seed.sql

e2e-test:
	TEST_MODE=true go test ./test/e2e/...

# etc.
fmt:
	go fmt ./...

model-gen:
	go run ./cmd/gentool/main.go
