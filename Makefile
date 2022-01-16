.DEFAULT_GOAL := build

build:
	@echo 'Building...'
	GOOS=linux GOARCH=amd64 go build -o=./bin/linux_amd64/ .
	GOOS=darwin GOARCH=amd64 go build -o=./bin/darwin_amd64/ .
	GOOS=darwin GOARCH=arm64 go build -o=./bin/darwin_arm64/ .

.PHONY: build