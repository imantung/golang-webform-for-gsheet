PROJ=$(shell pwd)

run:
	@go run ./cmd/golang-webform-for-gsheet

generate:
	@go install github.com/imantung/file_append
	@PROJ=${PROJ} go generate ./...
