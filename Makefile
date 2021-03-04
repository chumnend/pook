all: build serve

.PHONY: build
build:
	@echo "Building..."
	@cd ui/ && npm install && npm run build
	@mkdir -p bin/
	@cd bin/ && go build ../main.go
	@echo "Build complete."

.PHONY: serve
serve:
	@echo "Starting server..."
	@./bin/main

.PHONY: test
test:
	@echo "Running tests..."
	@go test

.PHONY: clean
clean:
	@echo "Cleaning binaries..."
	@rm -rf bin build ui/node_modules
	@echo "Clean complete."