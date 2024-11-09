.DEFAULT_GOAL := run

.PHONY: run
run:
	go run cmd/main.go

.PHONY: build
build:
	go build -o TiagoEcomm cmd/main.go

.PHONY: test
test:
	go test -v ./...

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: migration
migration:
	migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))	

.PHONY: migrate-up
migrate-up:
	go run cmd/migrate/main.go up

.PHONY: migrate-down
migrate-down:
	go run cmd/migrate/main.go down

.PHONY: lint
lint:
	golangci-lint run ./..