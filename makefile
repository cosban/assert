PROJECT := assert
PKG := "github.com/cosban/$(PROJECT)"
PKG_LIST := $(shell go list $(PKG)/... | grep -v /vendor/)

.PHONY: all dep build clean test

all: test race build

dep:
	@go get -t -v -d ./...

test: dep
	@go test -short $(PKG_LIST)

race: dep
	@go test -race -short $(PKG_LIST)

build: dep
	@go build -i -v $(PKG)

clean:
	@rm -f $(PROJECT)