build: format
	go build -v -buildvcs=false

test: format
	go test -v

format:
	@go fmt
