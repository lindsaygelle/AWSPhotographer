.PHONY: build

build:
	go mod download
	go mod tidy
	go mod verify
	go mod why
	@GOOS=linux GOARCH=amd64 go build -o lambda main.go
	zip lambda.zip lambda
test:
	go test ./...
