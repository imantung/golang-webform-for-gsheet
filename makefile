PROJ=$(shell pwd)

run:
	@go run ./cmd/golang-webform-for-gsheet

generate:
	@go install github.com/imantung/file_append
	@go install github.com/golang/mock/mockgen@v1.6.0
	@PROJ=${PROJ} go generate ./...

.PHONY: run generate