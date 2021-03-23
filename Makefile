all: build serve

.PHONY: build
build: build-react build-go

.PHONY: build-react
build-react:
	@echo "Building..."
	@cd web && npm run build
	@mv web/build .

.PHONY: build-go
build-go:
	@mkdir -p bin/
	@cd bin/ && go build ../cmd/pook/main.go

.PHONY: serve
serve:
	@echo "Starting server..."
	@./bin/main

.PHONY: clean
clean:
	@echo "Cleaning..."
	@rm -rf bin build
	@echo "Clean complete."