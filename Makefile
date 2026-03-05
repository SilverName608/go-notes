run:
	go run ./cmd/server/main.go

build:
	go build -o go-notes ./cmd/server/main.go

migrate-up:
	migrate -path ./migrations -database "postgres://postgres:4019@localhost:5432/go_notes?sslmode=disable" up

migrate-down:
	migrate -path ./migrations -database "postgres://postgres:4019@localhost:5432/go_notes?sslmode=disable" down