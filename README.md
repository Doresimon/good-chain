# good-chain

an experimental blockchain for public benefits.

## detail

- developping language: `golang`
- reference chain structure: `ethereum`
- db: `leveldb`
- communication: `rpc`
- dependency manager: `dep`

## Third Party Library

<!--
- `dep`

  > go get -u github.com/golang/dep/cmd/dep

  > dep ensure

- `leveldb`

  > go get -u github.com/syndtr/goleveldb/leveldb

- `go-libp2p`

  > go get -u github.com/libp2p/go-libp2p@v6.0.12

- `install dependencies`

  > dep ensure
-->

- `go mod`

  > export GO111MODULE=on

  > go mod init

  > go mod tidy

## Reference

1. [gobyexample](https://gobyexample.com)

## go doc

[![GoDoc](https://godoc.org/github.com/Doresimon/good-chain?status.svg)](https://godoc.org/github.com/Doresimon/good-chain)

## test

```bash
go test -v
```
