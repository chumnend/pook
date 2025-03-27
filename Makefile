include .env

all: build

.PHONY: build
build:
	@echo "Building..."
	@mkdir -p bin/
	@cd ./bin && go build ../cmd/main.go
	@echo "Build complete."

.PHONY: start
start:
	@echo "Executing..."
	@./bin/main

.PHONY: test
test:
	@echo "Running tests..."
	@go test ./internal/...

.PHONY: clean 
clean:
	@echo "Cleaning binaries..."
	@rm -rf bin
	@echo "Clean complete."

.PHONY: migrate
migrate:
	migrate -database "${PG_URL}?sslmode=disable" -path migrations up

.PHONY: rollback
rollback:
	migrate -database "${PG_URL}?sslmode=disable" -path migrations down
