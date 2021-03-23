all: build serve

.PHONY: build
build: build-react build-go

.PHONY: build-react
build-react:
	@echo "Building..."
	@if [ ! -d "web/node_modules" ]; then \
  	cd web && npm install; \
	fi
	@cd web && npm run build

.PHONY: build-go
build-go:
	@mkdir -p bin/
	@cd bin/ && go build ../main.go

.PHONY: serve
serve:
	@echo "Starting server..."
	@./bin/main

.PHONY: clean
clean:
	@echo "Cleaning..."
	@rm -rf bin web/build
	@echo "Clean complete."