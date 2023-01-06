server:
	go run main.go

test:
	go test -v -cover ./...

.PHONY: server test