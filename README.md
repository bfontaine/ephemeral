# ephemeral

`ephemeral` is a lightweight library to fire an ephemeral Web server. It’s
useful when you’re writing a CLI app which needs to authenticate against a
remote server using a callback URI.

[![GoDoc](https://godoc.org/github.com/bfontaine/ephemeral?status.svg)](https://godoc.org/github.com/bfontaine/ephemeral)
[![Build Status](https://travis-ci.org/bfontaine/ephemeral.svg?branch=master)](https://travis-ci.org/bfontaine/ephemeral)

## Usage

Ephemeral provides a small `http`-like API. Start by creating a server:

```go
s := ephemeral.New()
```

Then declare your handler like you’d do with a classic HTTP server, except that
it takes the server as its first argument:

```go
s.HandleFunc("/", func(s *ephemeral.Server,
    w http.ResponseWriter, r *http.Request) {

    w.Write([]byte("Ok bye :)\n"))

    s.Stop("foo")
})
```

The server exposes a `Stop` method which takes one argument that can be
whatever you want, including `nil`.

Start the server:

```go
s.Listen(":8080")
```

The method won’t return until the server is stopped. It returns the argument
you gave to the `.Stop()` call as well as an `error`.

## Install

    go get github.com/bfontaine/ephemeral
