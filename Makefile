clean:
	go clean -modcache

install:
	go install -v

fmt:
	gofmt -w .

.PHONY: clean install fmt

