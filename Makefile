BINARY := netbox-cli
IMAGE  := netbox-cli

.PHONY: build docker clean lint

build:
	go build -o $(BINARY) .

docker:
	docker build -t $(IMAGE) .

clean:
	rm -f $(BINARY)

lint:
	golangci-lint run ./...
