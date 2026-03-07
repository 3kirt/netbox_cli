BINARY := netbox-cli
IMAGE  := netbox-cli

.PHONY: build docker clean

build:
	go build -o $(BINARY) .

docker:
	docker build -t $(IMAGE) .

clean:
	rm -f $(BINARY)
