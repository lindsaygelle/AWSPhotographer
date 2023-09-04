.PHONY: build

build:
	@GOOS=linux GOARCH=amd64 go build -o lambda main.go
	zip lambda.zip lambda
