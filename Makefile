IMAGE=pongo2-runner
TAG=latest


build:
	mkdir -p build/
	CGO_ENABLED=0 go build -o build/pongo2-runner ./cmd/pongo2-runner

install:
	cp ./build/pongo2-runner ~/go/bin/

.PHONY: build install

docker-build:
	docker build \
		-t "$(IMAGE):$(TAG)" \
		.

docker-run:
	docker run \
		--rm \
		--entrypoint=sh \
		-it \
		--name pongo2-runner \
		"$(IMAGE):$(TAG)"