all: build serve

.PHONY: build
build:
	@if [ ! -d ui/node_modules ]; then\
		echo "Installing npm_modules" && cd ui/ && npm install;\
	fi
	@echo "Building..."
	@cd ui/ && npm run build
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
	@echo "Cleaning..."
	@rm -rf bin build ui/node_modules
	@echo "Clean complete."