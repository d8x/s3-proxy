BINARY_NAME=sgw
VERSION=$(shell git describe --tags --always --dirty)

.PHONY: build
build:
	@echo "Building..."
	@go build -o $(BINARY_NAME) -v -ldflags "-X main.version=$(VERSION)"
	@echo "Done, version: $(VERSION)"


.PHONY: test
test:
	@echo "Running tests..."
	@go test -v ./...
	@echo "Done"

.PHONY: docker-build-linux-x86
docker-build-linux-x86:
	@echo "Building docker image..."
	@docker buildx build --platform linux/amd64 -t ghcr.io/d8x/$(BINARY_NAME):$(VERSION) -t ghcr.io/d8x/$(BINARY_NAME):latest .
	@echo "Done, version: ghcr.io/d8x/$(BINARY_NAME):$(VERSION)"

.PHONY: docker-build-push-linux-x86
docker-build-push-linux-x86:
	@echo "Building docker image..."
	@docker buildx build --platform linux/amd64 -t ghcr.io/d8x/$(BINARY_NAME):$(VERSION) -t ghcr.io/d8x/$(BINARY_NAME):latest . --push
	@echo "Done, version: ghcr.io/d8x/$(BINARY_NAME):$(VERSION)"
