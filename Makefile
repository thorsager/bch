BINARY_NAME=bch

.PHONY: build
build:
	go build -o $(BINARY_NAME) -v

install:
	go install -v

clean:
	go clean