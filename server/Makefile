BINARY=pinbox-server

$(BINARY): test build

build:
	go build -o $(BINARY) cmd/pinbox-server/main.go

clean:
	rm -f $(BINARY)

test:
	go test -v -race -cover

