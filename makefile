# include .env
# export

.PHONY: run lint build clear

run:
	go mod tidy
	go run ./cmd/app/main.go

build:
	go mod tidy
	go build ./cmd/app/main.go
	./main

lint:
	golangci-lint run -c .golangci.yml
