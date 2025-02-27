# Makefile
.PHONY: build run test lint clean docker-up docker-down

build:
	go build -o bin/api cmd/api/main.go

run:
	go run cmd/api/main.go

test:
	go test -v ./...

lint:
	golangci-lint run

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

clean:
	rm -rf bin/