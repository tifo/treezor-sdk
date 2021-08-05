NAME	:= treezor-sdk

.DEFAULT: all

all: gen build test

build:
	$(info + [$(NAME)] $@)
	go build -v ./...

test:
	$(info + [$(NAME)] $@)
	go test -v -cover ./...

fmt:
	$(info + [$(NAME)] $@)
	golangci-lint run --fix --issues-exit-code=0 >/dev/null

check:
	$(info + [$(NAME)] $@)
	golangci-lint run

gen: treezor_accessors.go
treezor_accessors.go:
	$(info + [$(NAME)] gen)
	go generate .

.PHONY: all build test fmt check gen treezor_accessors.go
