.PHONY: build
build:
	go build -v ./cmd/master

.DEFAULT_GOAL := build