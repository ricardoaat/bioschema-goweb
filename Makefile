.PHONY: \
	build \
	windowsbuild \
	macbuild \

GO_PATH = $(shell echo $(GOPATH) | awk -F':' '{print $$1}')
GO_SRC = $(shell pwd | xargs dirname | xargs dirname | xargs dirname)
DEPLOY_PATH := build/
BIN_NAME := bioschemas-goweb
VERSION=$(shell git describe --tags)
BUILD=$(shell date +%FT%T%z)
LDFLAGS=-ldflags "-X main.version=${VERSION} -X main.buildDate=${BUILD}"

build:
	go build -i -v $(LDFLAGS) -o $(DEPLOY_PATH)example/$(BIN_NAME) main.go

windowsbuild:
	env GOOS=windows GOARCH=amd64 go build -i -v $(LDFLAGS) -o $(DEPLOY_PATH)$(BIN_NAME)_windows_64.exe main.go
	env GOOS=windows GOARCH=386 go build -i -v $(LDFLAGS) -o $(DEPLOY_PATH)$(BIN_NAME)_windows_x86.exe main.go

linuxbuild:
	env GOOS=linux GOARCH=amd64 go build -i -v $(LDFLAGS) -o $(DEPLOY_PATH)$(BIN_NAME)_linux_64 main.go
	env GOOS=linux GOARCH=386 go build -i -v $(LDFLAGS) -o $(DEPLOY_PATH)$(BIN_NAME)_linux_x86 main.go

macbuild:
	env GOOS=darwin GOARCH=amd64 go build -i -v $(LDFLAGS) -o $(DEPLOY_PATH)$(BIN_NAME)_mac_64 main.go
	env GOOS=darwin GOARCH=386 go build -i -v $(LDFLAGS) -o $(DEPLOY_PATH)$(BIN_NAME)_mac_x86 main.go

build-all: windowsbuild linuxbuild	macbuild