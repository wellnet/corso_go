include .env

.PHONY: run

default: run

run:
	@export API_KEY=$(API_KEY);go run main.go -port=8081
