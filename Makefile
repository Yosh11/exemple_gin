build:
	go build -v ./cmd/main.go

test:
	richgo test -v -cover -timeout 30s ./...
# for docker
migrate_up:
	migrate -path ./scheme -database 'postgres://postgres:qwerty@db:5432/postgres?sslmode=disable' up

migrate_down:
	migrate -path ./scheme -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' down

.DEFAULT_GOAL := build
