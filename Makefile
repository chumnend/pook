all: build serve

.PHONY: build
build: build-react build-go

.PHONY: build-react
build-react:
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
	@./bin/main

.PHONY: test
test:
	@go test -v
	@cd web && npm test -- --watchAll=false

.PHONY: clean
clean:
	@rm -rf bin web/build web/node_modules
	@echo "Clean complete."