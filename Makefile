NAME	:= treezor-sdk

.DEFAULT: all

all: gen build test

build: gen
	$(info + [$(NAME)] build)
	go build -v ./...

test:
	$(info + [$(NAME)] test)
	go test -v ./...

gen: go-generate
go-generate:
	$(info + [$(NAME)] go-generate)
	go generate .

.PHONY: all build test gen go-generate
