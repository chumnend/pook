.PHONY: start-client
start-client:
	go run app/pook-client/main.go

.PHONY: start-api
start-api:
	go run app/pook-api/main.go
