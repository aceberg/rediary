mod:
	rm go.mod || true && \
	rm go.sum || true && \
	go mod init github.com/aceberg/red && \
	go mod tidy

run:
	cd cmd/red/ && \
	go run .

fmt:
	go fmt ./...

lint:
	golangci-lint run
	golint ./...

check: fmt lint

go-build:
	cd cmd/red/ && \
	CGO_ENABLED=0 go build -o ../../tmp/red .