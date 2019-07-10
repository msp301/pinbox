BINARY=pinbox-server

.PHONY: build clean test

$(BINARY): test build

build:
	go get -d ./...
	go build -o $(BINARY) cmd/pinbox-server/main.go

clean:
	rm -f $(BINARY)

test:
	go test -v -race -cover

