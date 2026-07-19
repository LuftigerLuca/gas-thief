.PHONY: build lint fmt fmt-check test check

build:
	go build ./...

lint:
	golangci-lint run ./...

fmt:
	gofmt -w -s .

fmt-check:
	@[ -z "$$(gofmt -l -s .)" ] || (echo "Files not formatted:" && gofmt -l -s . && exit 1)

test:
	go test ./...

check: fmt-check lint build test
