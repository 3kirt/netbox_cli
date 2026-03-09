BINARY  := netbox-cli
IMAGE   := netbox-cli
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)

.PHONY: build docker clean lint

build:
	go build -ldflags "-X github.com/kirtis/netbox-cli/cmd.version=$(VERSION)" -o $(BINARY) .

docker:
	docker build -t $(IMAGE) .

clean:
	rm -f $(BINARY)

lint:
	golangci-lint run ./...
