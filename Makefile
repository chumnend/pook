all: build serve

# Build React and Go assets in bin folder
.PHONY: build
build: build-react build-go

# Build React assets in bin folder
.PHONY: build-react
build-react:
	@if [ ! -d "client/node_modules" ]; then \
  	cd client && npm install; \
	fi
	@cd client && npm run build
	@echo "React files built."

# Build Go assets in bin folder
.PHONY: build-go
build-go:
	@mkdir -p bin/
	@cd bin/ && go build ../server/cmd/pook/main.go
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
	@cd client && npm test -- --watchAll=false

# Executes tests for Go packages
.PHONY: test-go
test-go:
	@if [ ! -d "client/build" ]; then \
  	cd client && npm run build; \
	fi
	@go test ./server/tests/...

# Executes only unit tests for Go packages
.PHONY: unittest
unittest:
	@if [ ! -d "client/build" ]; then \
  	cd client && npm run build; \
	fi
	@go test -short ./server/tests/...

# Cleans up assets and node_modules
.PHONY: clean
clean:
	@rm -rf bin client/build client/node_modules
	@echo "Clean complete."