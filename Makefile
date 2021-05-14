all: build serve

.PHONY: build
build: build-react build-go

.PHONY: build-react
build-react:
	@if [ ! -d "web/node_modules" ]; then \
  	cd web && npm install; \
	fi
	@cd web && npm run build
	@echo "React files built."

.PHONY: build-go
build-go:
	@mkdir -p bin/
	@cd bin/ && go build ../main.go
	@echo "Go files built."

.PHONY: serve
serve:
	@./bin/main

.PHONY: test
test:
	@if [ ! -d "web/build" ]; then \
  	cd web && npm run build; \
	fi
	@go test -v
	@cd web && npm test -- --watchAll=false

.PHONY: clean
clean:
	@rm -rf bin web/build web/node_modules
	@echo "Clean complete."