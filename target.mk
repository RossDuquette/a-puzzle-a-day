build: format
	go build -v -buildvcs=false

run: build
	./a-puzzle-a-day

test: format
	go test -v ./...

format:
	@go fmt ./...
