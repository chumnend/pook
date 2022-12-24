.PHONY: run
run:
	@go mod tidy && go mod download && go run ./cmd/app

.PHONY: test
test: 
	@go test -cover -covermode=atomic ./internal/...