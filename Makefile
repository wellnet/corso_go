include .env

.PHONY: run

default: run

run:
	@go run main.go -port=8081
