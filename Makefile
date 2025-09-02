build:
	@go build -o bin/planning-poker-backend ./cmd/server

test:
	@go test -v ./...

run: build
	@echo "Running the server..."
	@./bin/planning-poker-backend

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@, $(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down

lint:
	@golangci-lint run ./... --config=.golangci.yml --fix

gen-docs:
	@swag init -g ./api/main.go -d cmd,./internal && swag fmt

test-store:
	@go test -v ./internal/store
