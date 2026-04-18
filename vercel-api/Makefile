.PHONY: build run

OS := $(shell uname -s 2>/dev/null | tr '[:upper:]' '[:lower:]' || echo darwin)

build:
	GOOS=$(OS) go build -o bin/api ./cmd/api

run: build
	./bin/api
