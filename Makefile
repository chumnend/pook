.PHONY: build
build:
	@echo "Building..."
	@mkdir -p bin/bookings
	@cd ./bin/bookings && go build ../../cmd/bookings/main.go
	@echo "Build complete."

.PHONY: start
start:
	@echo "Starting server..."
	@go run cmd/bookings/main.go

.PHONY: clean
clean:
	@echo "Cleaning binaries..."
	@rm -rf bin
	@echo "Clean complete."