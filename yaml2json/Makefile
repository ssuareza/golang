dev:
	docker build -t yaml2json .
	docker run -it --rm -v ./:/app yaml2json

sh:
	docker build -t yaml2json .
	docker run -it --rm -v ./:/app yaml2json sh

version=$(shell git describe --all --dirty --long | awk -F"-|/" '/^heads/ {print $$2 "-" substr($$4, 2) "-" $$5}; /^tags/ { print $$2 }')
build_args=-ldflags "-X main.versionString=$(version)" ./cmd/yaml2json
files=$(shell find cmd -type f)

build: build-linux build-darwin

build-linux: build/yaml2json-$(version)-linux-amd64

build/yaml2json-$(version)-linux-amd64: ${files}
	GOARCH=amd64 GOOS=linux go build -o $@ $(build_args)

build-darwin: build/yaml2json-$(version)-darwin-amd64
build/yaml2json-$(version)-darwin-amd64: ${files}
	GOARCH=amd64 GOOS=darwin go build -o $@ $(build_args)
