lint:
	golangci-lint run

test:
	go test ./...

format:
	gofmt -w .

.PHONY: lint test format