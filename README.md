# wikiproxy

A PoC Go app that acts as a proxy for Wikipedia, similar to wikiless but written in Go

Known issues:

* Serves images directly from wikimedia still
* Language is hardcoded to en
* No search

## Usage

```shell
# Run from source
go run proxy.go

# Run compiled version on port 8123
PORT=8123 ./wikiproxy
```

Then browse to eg. http://localhost:8000