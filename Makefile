.PHONY: start-dev
start-dev:
	go run app/pook-api/main.go

.PHONY: serve
serve:
	go run app/pook-client/main.go