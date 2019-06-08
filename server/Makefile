BINARY=pinbox-server

$(BINARY): test
	go build -o $(BINARY) cmd/pinbox-server/main.go

clean:
	rm -f $(BINARY)

test:
	go test -v -race -cover

