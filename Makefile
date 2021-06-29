all: build serve

# Build React and Go assets in bin folder
.PHONY: build
build: build-react build-go

# Build React assets in bin folder
.PHONY: build-react
build-react:
	@if [ ! -d "web/node_modules" ]; then \
  	cd web && npm install; \
	fi
	@cd web && npm run build
	@echo "React files built."

# Build Go assets in bin folder
.PHONY: build-go
build-go:
	@mkdir -p bin/
	@cd bin/ && go build ../cmd/pook/main.go
	@echo "Go files built."

# Starts the app on port provided in .env file
.PHONY: serve
serve:
	@./bin/main

# Executes tests for Go packages and React app
.PHONY: test
test: test-react test-go

# Executes tests for React app
.PHONY: test-react
test-react:
	@cd web && npm test -- --watchAll=false

# Executes tests for Go packages
.PHONY: test-go
test-go:
	@if [ ! -d "web/build" ]; then \
  	cd web && npm run build; \
	fi
	@go test ./internal/... -cover -covermode=atomic

# Cleans up assets and node_modules
.PHONY: clean
clean:
	@rm -rf bin web/build web/node_modules
	@echo "Clean complete."