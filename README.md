# wolgo
Wake on LAN


## one binary build for linux


```bash
GOOS=linux CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo --ldflags '-extldflags "-static"' -o build/wakeonlan cli/cli.go
```
