package simple_fasthttp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/valyala/fasthttp"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func startTestServerError() *fasthttp.Server {
	return &fasthttp.Server{
		Handler: func(ctx *fasthttp.RequestCtx) {
			if string(ctx.Path()) == "/ping" {
				ctx.SetStatusCode(400)
				res, _ := json.Marshal(ErrorResponse{Status: 400, Message: "error"})
				ctx.SetBody(res)
			} else {
				ctx.SetStatusCode(404)
			}
		},
	}
}

func TestDoError(t *testing.T) {
	srv := startTestServerError()
	go srv.ListenAndServe(":8085")
	defer srv.Shutdown()

	u, _ := url.Parse("http://localhost:8085/ping")
	opts := RequestOptions{
		URL:     u,
		Timeout: 2 * time.Second,
	}

	_, httpData, err := Do[any, ErrorResponse](context.Background(), opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if httpData.Status() != 400 {
		t.Fatalf("status code not 404: %d", httpData.Status())
	}
	fmt.Println(httpData.Payload())
}
