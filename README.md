# micro-go

micro service framework in Go

## Features

* Monorepo - Sharing code between microservies
* [PGV](https://github.com/envoyproxy/protoc-gen-validate) - input validation
* [koanf](https://github.com/knadh/koanf) - env, config
* [logr](https://github.com/go-logr/logr) - log interface
* [zap](https://github.com/uber-go/zap) - log implementation
* grpc
* [req](https://github.com/imroc/req) - http client
* [errors](https://github.com/cockroachdb/errors) - error handling
* [ent](https://entgo.io/) - Graph-Based ORM
* [wire](https://github.com/google/wire) - dependency injection
* multi-stage-multi-target Dockerfile
* [ko](https://github.com/google/ko) - One Step build/publish/deploy
* [govvv](https://github.com/ahmetb/govvv)
* [golangci]
* [buf]
* [super-linter](https://github.com/github/super-linter)
* github actions
* [tableflip](https://github.com/cloudflare/tableflip) - graceful process restart
* graphql: [rejoiner](https://github.com/google/rejoiner) [gqlgen](https://gqlgen.com/)

## third-party tools

```sh
# github CLI
brew install hub
# for mac, use brew to install protobuf
brew install protobuf
# VS Code plugin `vscode-proto3` need clang-format
brew install clang-format
# k8s tool similar to helm  (optional)
# generate fill k8s yaml files from overlays
brew install kustomize
# kubeval - validate one or more Kubernetes config files(optional)
brew tap instrumenta/instrumenta
brew install kubeval
# Manage Your lk8s In Style!
brew install derailed/k9s/k9s
# grpc cli client (optional)
brew install grpcurl
# bloomrpc is a UI client for gRPC (optional)
# install `bloomrpc` via `brew` into ~/Applications)
brew cask install --appdir=~/Applications bloomrpc
# gRPC mock server for testing
yarn global add bloomrpc-mock
# CHANGELOG generator
brew tap git-chglog/git-chglog
brew install git-chglog
# buf: proto tool https://buf.build/docs/tour-1
brew tap bufbuild/buf
brew install buf
# git flow
brew install git-flow-avh
```

golang tools

```sh
# go better build tool
go install github.com/ahmetb/govvv@latest
# for static check/linter
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
# linter and tool for proto files
# *** (if you use brew to install buf, skip next line) ***
go install github.com/bufbuild/buf/cmd/buf@latest
# kind - kubernetes in docker (optional)
go install sigs.k8s.io/kind@latest
# go lang  build/publish/deploy tool (optional)
go install github.com/google/ko/cmd/ko@latest

# fetch protoc plugins into $GOPATH
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# to add tags to go struct
go install github.com/srikrsna/protoc-gen-gotag@latest

# Installing PGV can currently only be done from source: 
# from user's home directory, run
go get -d github.com/envoyproxy/protoc-gen-validate
cd ~/go/src/github.com/envoyproxy/protoc-gen-validate
git pull
make build

# goup checks if there are any updates for imports in your module.
# the main purpose is using it as a linter in continuous integration or in development process.
# Usage: goup -v -m ./...
go install github.com/rvflash/goup@latest
```

## Boilerplate codes

* [Makego](https://github.com/bufbuild/makego)
