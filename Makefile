all: build serve

.PHONY: build
build:
	@echo "Building..."
	@cd client && npm run build
	@mkdir -p bin/
	@cd bin/ && go build ../main.go
	@echo "Build complete."

.PHONY: serve
serve:
	@echo "Starting server..."
	@./bin/main

.PHONY: start
start:
	@echo "Starting server..."
	@go run main.go

.PHONY: clean
clean:
	@echo "Cleaning..."
	@rm -rf bin client/build
	@echo "Clean complete."