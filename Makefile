ifneq (,$(wildcard .env))
    include .env
    export
endif

run:
	go run cmd/server/main.go

up:
	docker-compose up -d

down:
	docker-compose down


migrate-up:
	goose -dir migrations postgres "$(DATABASE_URL)" up

migrate-down:
	goose -dir migrations postgres "$(DATABASE_URL)" down

migrate-create:
	goose -dir migrations create $(name) sql

test:
	go test ./... -cover

.PHONY: build run start

start:
	go build -o bin/app ./cmd/server && ./bin/app