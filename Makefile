# This is a generic Makefile that can be used for go projects to do building, linting, and static analysis
SHELL=/bin/bash
GO111MODULE=on

.PHONY: install check build

.PHONY:
install:
	@go get golang.org/x/tools/cmd/goimports
	@go get golang.org/x/lint/golint
	@go get honnef.co/go/tools/...

.PHONY:
build:
	@go build ./...

.PHONY:
check: fmt lint vet imports staticcheck

.PHONY:
fmt:
	@go fmt ./...

.PHONY:
lint:
	@golint ./...

.PHONY:
vet:
	@go vet ./...

.PHONY:
imports:
	@goimports -w .

.PHONY:
staticcheck:
	@staticcheck ./...


#@if test -s "main.go"; then go build; else echo "No main.go file found"; fi;
