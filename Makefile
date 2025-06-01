include .env

all: build start

.PHONY: build
build:
	@echo "Building..."
	@mkdir -p bin/
	@cd ./bin && go build ../cmd/main.go
	@if [ ! -d "./web/node_modules" ]; then \
			echo "Installing dependencies..."; \
			cd ./web && yarn; \
	fi
	@cd ./web && yarn build
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
	@rm -rf bin web/**/dist web/**/node_modules
	@echo "Clean complete."

.PHONY: migrate
migrate:
	migrate -database "${PG_URL}?sslmode=disable" -path migrations up

.PHONY: rollback
rollback:
	migrate -database "${PG_URL}?sslmode=disable" -path migrations down
