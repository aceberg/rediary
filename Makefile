mod:
	rm go.mod || true && \
	rm go.sum || true && \
	go mod init github.com/aceberg/rediary && \
	go mod tidy

run:
	cd cmd/rediary/ && \
	go run .

fmt:
	go fmt ./...

lint:
	golangci-lint run
	golint ./...

check: fmt lint

go-build:
	cd cmd/rediary/ && \
	CGO_ENABLED=0 go build -o ../../tmp/rediary .