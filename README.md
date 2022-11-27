# klev API client library for Golang

A wrapper around [klev](https://klev.dev) API that makes it easier to use it in Golang.

## Installation

Installation is as simple as using `go get`:

```
go get github.com/klev-dev/klev-api-go
```

## Documentation

You can find documentation at:
 * Full [API Reference](https://klev.dev/api)
 * Golang [package docs](https://pkg.go.dev/github.com/klev-dev/klev-api-go)

## Quickstart

See the [examples repository](https://github.com/klev-dev/klev-examples) for additional examples.

### klev hello world

```go
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	klev "github.com/klev-dev/klev-api-go"
)

func main() {
	fmt.Println(hello(context.Background()))
}

func hello(ctx context.Context) error {
	cfg := klev.NewConfig(os.Getenv("KLEV_TOKEN_DEMO"))
	client := klev.New(cfg)

	log, err := client.LogCreate(ctx, klev.LogIn{})
	if err != nil {
		return err
	}

	_, err = client.Post(ctx, log.LogID, time.Time{}, nil, []byte("hello world!"))
	if err != nil {
		return err
	}

	msg, err := client.Get(ctx, log.LogID, 0)
	if err != nil {
		return err
	}
	fmt.Println(string(msg.Value))

	return nil
}
```
