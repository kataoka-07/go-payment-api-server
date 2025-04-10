.PHONY: up stop clean fmt run

# Docker DB
up:
	docker compose up --build -d

stop:
	docker compose stop

clean:
	docker compose down

fmt:
	go fmt ./...

# Server
run:
	go run ./cmd/payment-server/main.go

model-gen:
	go run ./cmd/gentool/main.go