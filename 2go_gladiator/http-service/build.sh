#!/bin/bash
go mod tidy
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/linux/http-server cmd/http-server/main.go
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/linux/http-client cmd/http-client/main.go
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/linux/cert-retriever cmd/cert-retriever/main.go

