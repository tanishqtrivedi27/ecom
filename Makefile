build:
	@go build -o bin/ecom cmd/main.go

test:
	@go test -v ./..

run: build
	@./bin/ecom

migrate:
	@go run cmd/migrate/main.go
