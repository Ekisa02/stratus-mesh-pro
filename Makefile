BINARY=bin/stratus-mesh

.PHONY: build run test fmt vet

build:
	go build -o $(BINARY) ./cmd/stratus-mesh

run: build
	./$(BINARY)

test:
	go test ./... -v

fmt:
	gofmt -s -w .

vet:
	go vet ./...
