build:
	@go build -o bin/reurb ./cmd

run: build
	@./bin/reurb

stop: down

up:
	@docker compose up -d

down:
	@docker compose down

test:
	@go test -v ./...
