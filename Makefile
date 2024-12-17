build:
	@go build -o bin/basic-goolang-api ./cmd/migrations/main.go

test:
	@go test -v ./...

run: build
	@./bin/basic-goolang-api
