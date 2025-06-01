# Makefile for web-server


run:
	go run cmd/server.go

test:
	go test ./... -v
