# wikiproxy

A PoC Go app that acts as a privacy-protecting proxy for Wikipedia, similar to [wikiless](https://wikiless.org/) but written in Go

This serves just Wikipedia pages, no tracking, no ads, no search (yet)

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