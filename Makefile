BINARY_NAME=go-rest-cli

build:
	go build -o bin/${BINARY_NAME} main.go

clean:
	go clean
	rm -rf ./bin

lint:
	golangci-lint run --enable-all
