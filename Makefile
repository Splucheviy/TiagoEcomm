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