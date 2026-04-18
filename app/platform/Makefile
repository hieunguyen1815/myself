.PHONY: build-profile run-profile start.devstack stop.devstack restart.devstack

OS := $(shell uname -s 2>/dev/null | tr '[:upper:]' '[:lower:]' || echo darwin)

build-profile:
	GOOS=$(OS) go build -o bin/profile ./modules/profile/cmd/profile

run-profile: build-profile
	./bin/profile

start.devstack:
	docker compose -f docker-compose.yaml up --build -d

stop.devstack:
	docker compose -f docker-compose.yaml down

restart.devstack: stop.devstack start.devstack
