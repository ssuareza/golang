dev:
	docker build -t yaml2json .
	docker run -it --rm -v ./:/app yaml2json

sh:
	docker build -t yaml2json .
	docker run -it --rm -v ./:/app yaml2json sh

build_args=-ldflags "-X main.versionString=$(version)" ./cmd/yaml2json

build: build-linux build-darwin

build-linux: build/yaml2json-$(version)-linux-amd64

build/yaml2json-$(version)-linux-amd64: ${files}
	GOARCH=amd64 GOOS=linux go build -o $@ $(build_args)

build-darwin: build/yaml2json-$(version)-darwin-amd64
build/yaml2json-$(version)-darwin-amd64: ${files}
	GOARCH=amd64 GOOS=darwin go build -o $@ $(build_args)
