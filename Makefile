PHONY: generate
generate:
	buf generate --path=./proto/url_shorter.proto



PHONY: test
test:
	go test ./...
