.PHONY: install-all
install-all: install-go install-react

.PHONY: install-go
install-go:
	@go mod tidy && go mod download

.PHONY: run-server
run-server:
	@go run ./cmd/app

.PHONY: test-server
test-server: 
	@go test -cover -covermode=atomic ./internal/...

.PHONY: install-react
install-react:
	@cd web/pook-react && npm install

.PHONY: build-client
build-client:
	@cd web/pook-react/ && npm run build

.PHONY: run-client
run-client:
	@go run ./cmd/web