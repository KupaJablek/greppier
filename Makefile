BINARY_NAME=greppier.bin

build:
	go build -o bin/${BINARY_NAME} .

clean:
	go clean
	rm ./bin/${BINARY_NAME}
