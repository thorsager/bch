BINARY_NAME=bch

.PHONY: build
build:
	go build -o $(BINARY_NAME) -v

clean:
	go clean