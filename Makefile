.PHONY: all
all: help

.PHONY: genpb # generate Golang protobuf files
genpb:
	buf generate

.PHONY: build-app # go build app
build-app:
	go build ./cmd/app

.PHONY: lint # golangci-lint
lint:
	golangci-lint run

.PHONY: help # Generate list of targets with descriptions
help:
	@grep '^.PHONY: .* #' Makefile | sed 's/\.PHONY: \(.*\) # \(.*\)/\1	\2/' | expand -t20
