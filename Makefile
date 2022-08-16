IMAGE=pongo2-runner
TAG=latest

IMAGE_BUILD=$(IMAGE)-build

build:
	mkdir -p build/
	CGO_ENABLED=0 go build -o build/pongo2-runner ./cmd/pongo2-runner

build-linux:
	mkdir -p build/
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/pongo2-runner-linux-x86_64 ./cmd/pongo2-runner
	strip build/pongo2-runner-linux-x86_64

build-darwin:
	mkdir -p build/
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/pongo2-runner-darwin-x86_64 ./cmd/pongo2-runner

build-windows:
	mkdir -p build/
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/pongo2-runner-windows-x86_64.exe ./cmd/pongo2-runner

build-all: build-linux build-darwin build-windows

install:
	cp ./build/pongo2-runner ~/go/bin/

.PHONY: build install

docker-build:
	docker build \
		-t "$(IMAGE):$(TAG)" \
		.

docker-build-all:
	docker build \
		-t "$(IMAGE_BUILD):$(TAG)" \
		-f Dockerfile.build \
		.
	docker run --rm \
		-d \
		--entrypoint=tail \
		--name pongo2-build-result \
		"$(IMAGE_BUILD):$(TAG)" \
		-f /dev/null

	docker cp pongo2-build-result:/app/build ./
	docker rm -f pongo2-build-result

docker-run:
	docker run \
		--rm \
		--entrypoint=sh \
		-it \
		--name pongo2-runner \
		"$(IMAGE):$(TAG)"