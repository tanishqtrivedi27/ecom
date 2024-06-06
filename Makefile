build:
	@go build -o bin/ecom cmd/main.go

test:
	@go test -v ./..

migrate:
	@go run cmd/migrate/main.go

run: migrate build
	@./bin/ecom
