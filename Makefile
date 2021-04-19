test:
	go test ./...

gen:
	go generate ./...

dev:
	go run cmd/fiber/main.go

build:
	go build -o cmd/fiber/fiber cmd/fiber/main.go

run:
	cmd/fiber/fiber

setup:
	docker-compose -f development/docker-compose.yml up -d
