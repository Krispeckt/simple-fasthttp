# 🎈 simple-fasthttp

`simple-fasthttp` is a lightweight wrapper around
[fasthttp](https://github.com/valyala/fasthttp) for making HTTP requests
in Go with a clean and type-safe API.\
It provides a simple interface for GET/POST requests, query parameters,
headers, request bodies, timeouts, and JSON unmarshalling into typed
responses.

------------------------------------------------------------------------

## Features

- ⚡ Built on top of **fasthttp** (high performance)
- ✅ Type-safe response decoding with generics
- ⏳ Configurable timeouts
- 🔑 Support for headers, query params, and request body
- 📦 Ready-to-use shortcuts for common HTTP methods

------------------------------------------------------------------------

## Installation

``` bash
go get github.com/krispeckt/simple-fasthttp
```

Make sure you have Go 1.18+ installed (generics are used).

------------------------------------------------------------------------

## Usage

### Simple GET usage

``` go
package main

import (
    "context"
    "fmt"
    "net/url"
    "time"

    shttp "github.com/krispeckt/simple-fasthttp"
)

type Response struct {
    Message string `json:"message"`
}

func main() {
    ctx := context.Background()
    u, _ := url.Parse("http://localhost:8080/ping")

    res, _, err := shttp.GetSimple[Response](ctx, u)
    if err != nil {
        panic(err)
    }

    fmt.Println("Response:", res.Message)
}
```

### Basic GET request

``` go
package main

import (
    "context"
    "fmt"
    "net/url"
    "time"

    shttp "github.com/krispeckt/simple-fasthttp"
)

type Response struct {
    Message string `json:"message"`
}

func main() {
    ctx := context.Background()
    u, _ := url.Parse("http://localhost:8080/ping")

    opts := shttp.RequestOptions{
        URL:     u,
        Timeout: 3 * time.Second,
    }

    res, _, err := shttp.Do[Response, any](ctx, opts)
    if err != nil {
        panic(err)
    }

    fmt.Println("Response:", res.Message)
}
```

### POST request with JSON body

``` go
type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type LoginResponse struct {
    Token string `json:"token"`
}

func login() {
    ctx := context.Background()
    u, _ := url.Parse("http://localhost:8080/login")

    opts := shttp.RequestOptions{
        Method: "POST",
        URL:    u,
        Body: LoginRequest{
            Username: "admin",
            Password: "secret",
        },
        Headers: map[string]string{
            "Content-Type": "application/json",
        },
    }

    res, _, err := shttp.Do[LoginResponse, any](ctx, opts)
    if err != nil {
        panic(err)
    }

    fmt.Println("Token:", res.Token)
}
```

------------------------------------------------------------------------

## API

### `RequestOptions`

``` go
type RequestOptions struct {
    Method  string
    URL     *url.URL
    Params  url.Values
    Headers map[string]string
    Body    any
    Timeout time.Duration
}
```

- **Method**: HTTP method (default: GET)
- **URL**: Target URL (required)
- **Params**: Query parameters
- **Headers**: Custom headers
- **Body**: JSON-encoded request body
- **Timeout**: Timeout for the request (default: 5s)

### `Do[T, E]`

``` go
func Do[T any, E any](ctx context.Context, opts RequestOptions) (*T, Http, error)
```

- `T` - success response type (JSON will be unmarshalled into this)
- `E` - error response type in http data

Returns: - `*T` - parsed response - `Http` - custom http type with http data - `error` - network/serialization errors

------------------------------------------------------------------------

## Running Tests

``` bash
go test ./...
```

------------------------------------------------------------------------

## License

MIT License. See [LICENSE](LICENSE) for details.
