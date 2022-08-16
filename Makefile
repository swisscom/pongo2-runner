build:
	mkdir -p build/
	CGO_ENABLED=0 go build -o build/pongo2-runner ./cmd/pongo2-runner

install:
	cp ./build/pongo2-runner ~/go/bin/

.PHONY: build install