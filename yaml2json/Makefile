app=yaml2json
version=$(shell git describe --all --dirty --long | awk -F"-|/" '/^heads/ {print $$2 "-" substr($$4, 2) "-" $$5}; /^tags/ { print $$2 }')
build_args=-ldflags "-X main.versionString=$(version)" ./cmd/$(app)
files=$(shell find cmd -type f)

dev:
	docker build -t $(app) .
	docker run -it --rm -v ./:/app $(app)

sh:
	docker build -t $(app) .
	docker run -it --rm -v ./:/app $(app) sh

build: build-linux build-darwin

build-linux: build/$(app)-$(version)-linux-amd64

build/$(app)-$(version)-linux-amd64: ${files}
	GOARCH=amd64 GOOS=linux go build -o $@ $(build_args)

build-darwin: build/$(app)-$(version)-darwin-amd64
build/$(app)-$(version)-darwin-amd64: ${files}
	GOARCH=amd64 GOOS=darwin go build -o $@ $(build_args)
