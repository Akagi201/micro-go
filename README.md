# micro-go

[![ci](https://github.com/Akagi201/micro-go/actions/workflows/ci.yml/badge.svg)](https://github.com/Akagi201/micro-go/actions/workflows/ci.yml)

micro service framework in Go

## Features

* Monorepo - Sharing code between microservies
* [PGV](https://github.com/envoyproxy/protoc-gen-validate) - input validation
* [cli](https://cli.urfave.org/) - urfave/cli
* [koanf](https://github.com/knadh/koanf) - env, config
* [logr](https://github.com/go-logr/logr) - log interface
* [zap](https://github.com/uber-go/zap) - log implementation
* [grpc](google.golang.org/grpc) - use pure GRPC
* [req](https://github.com/imroc/req) - http client
* [errors](https://github.com/cockroachdb/errors) - error handling
* [bun](https://github.com/uptrace/bun) - gorm
* [sqlc](https://github.com/kyleconroy/sqlc) - no gorm, generate code from sql schema
* [sqlx](https://github.com/jmoiron/sqlx) - no gorm
* [sql-migrate](https://github.com/rubenv/sql-migrate) - sql migrate
* [wire](https://github.com/google/wire) - dependency injection
* [govvv](https://github.com/ahmetb/govvv)
* [golangci](https://golangci-lint.run/) - lint tools
* [buf](https://buf.build) - protobuf management
* [super-linter](https://github.com/github/super-linter) - lint everything
* [tableflip](https://github.com/cloudflare/tableflip) - graceful process restart
* graphql: [rejoiner](https://github.com/google/rejoiner) [gqlgen](https://gqlgen.com/)

## Libs

* <https://github.com/destel/rill> - parallel processing like rayon in rust

## third-party tools

* <https://github.com/nikolaydubina/go-recipes>

```sh
# k8s tool similar to helm  (optional)
# generate fill k8s yaml files from overlays
brew install kustomize
# kubeval - validate one or more Kubernetes config files(optional)
brew tap instrumenta/instrumenta
brew install kubeval
# Manage Your lk8s In Style!
brew install derailed/k9s/k9s
# buf: proto tool https://buf.build/docs/tour-1
brew tap bufbuild/buf
brew install buf
```

golang tools

```sh
# go better build tool
go install github.com/ahmetb/govvv@latest
# for static check/linter
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# fetch protoc plugins into $GOPATH
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# to add tags to go struct
go install github.com/srikrsna/protoc-gen-gotag@latest

# goup checks if there are any updates for imports in your module.
# the main purpose is using it as a linter in continuous integration or in development process.
# Usage: goup -v -m ./...
go install github.com/rvflash/goup@latest
```

## Docs

* go mod update all: <https://gosamples.dev/update-all-packages/>

## After upgrade go version

```sh
go install -v golang.org/x/tools/...@latest
go install -v google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install -v google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install -v github.com/mgechev/revive@latest
go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```
