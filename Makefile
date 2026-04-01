build:
	CGO_ENABLED=0 go build -o drifter ./cmd/drifter/

run: build
	./drifter

test:
	go test ./...

clean:
	rm -f drifter

.PHONY: build run test clean
