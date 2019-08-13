# wolgo
[![GoDoc](https://godoc.org/github.com/minami14/wolgo?status.svg)](https://godoc.org/github.com/minami14/wolgo)
[![CircleCI](https://circleci.com/gh/minami14/wolgo.svg?style=shield)](https://circleci.com/gh/minami14/wolgo)

Wake on LAN


## one binary build for linux


```bash
GOOS=linux CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo --ldflags '-extldflags "-static"' -o build/wakeonlan cli/cli.go
```
