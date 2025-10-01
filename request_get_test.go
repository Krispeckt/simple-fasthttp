package simple_fasthttp

import (
	"context"
	"encoding/json"
	"net/url"
	"testing"
	"time"

	"github.com/valyala/fasthttp"
)

type PingResponse struct {
	Message string `json:"message"`
}

func startTestServerGet() *fasthttp.Server {
	return &fasthttp.Server{
		Handler: func(ctx *fasthttp.RequestCtx) {
			if string(ctx.Path()) == "/ping" {
				ctx.SetStatusCode(200)
				res, _ := json.Marshal(PingResponse{Message: "pong"})
				ctx.SetBody(res)
			} else {
				ctx.SetStatusCode(404)
			}
		},
	}
}

func TestDoGetSuccess(t *testing.T) {
	srv := startTestServerGet()
	go srv.ListenAndServe(":8085")
	defer srv.Shutdown()

	u, _ := url.Parse("http://localhost:8085/ping")
	opts := RequestOptions{
		URL:     u,
		Timeout: 2 * time.Second,
	}

	res, httpData, err := Do[PingResponse, Http[any]](context.Background(), opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if httpData.Status() != 200 {
		t.Fatalf("status code not 200: %d", httpData.Status())
	}
	if res.Message != "pong" {
		t.Errorf("expected pong, got %s", res.Message)
	}
}
