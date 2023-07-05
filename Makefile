run:
	go run ./cmd/potato/main.go

.PHONY: build
build:
	go build -o ./build ./cmd/...
