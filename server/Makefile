BINARY=pinbox-server

.PHONY: build clean test

$(BINARY): build test

build:
	go get -d ./...
	go build -o $(BINARY) cmd/pinbox-server/main.go

clean:
	rm -f $(BINARY)

test:
	go test -v -race -cover

