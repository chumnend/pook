all: build serve

.PHONY: build
build:
	@echo "Building..."
	@mkdir -p bin/
	@cd bin/ && go build ../main.go
	@echo "Build complete."

.PHONY: serve
serve:
	@echo "Starting server..."
	@./bin/main

.PHONY: clean
clean:
	@echo "Cleaning..."
	@rm -rf bin
	@echo "Clean complete."