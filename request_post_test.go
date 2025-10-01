package simple_fasthttp

import (
	"context"
	"encoding/json"
	"net/url"
	"testing"

	"github.com/valyala/fasthttp"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// mock POST server
func startTestServerPost() *fasthttp.Server {
	return &fasthttp.Server{
		Handler: func(ctx *fasthttp.RequestCtx) {
			if string(ctx.Path()) == "/login" {
				var req LoginRequest
				json.Unmarshal(ctx.PostBody(), &req)
				if req.Username == "admin" && req.Password == "secret" {
					ctx.SetStatusCode(200)
					res, _ := json.Marshal(LoginResponse{Token: "xyz123"})
					ctx.SetBody(res)
				} else {
					ctx.SetStatusCode(401)
				}
			}
		},
	}
}

func TestDoPostSuccess(t *testing.T) {
	srv := startTestServerPost()
	go srv.ListenAndServe(":8086")
	defer srv.Shutdown()

	u, _ := url.Parse("http://localhost:8086/login")
	opts := RequestOptions{
		Method: "POST",
		URL:    u,
		Body: LoginRequest{
			Username: "admin",
			Password: "secret",
		},
	}

	res, httpData, err := Do[LoginResponse, Http[any]](context.Background(), opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if httpData.Status() != 200 {
		t.Fatalf("status code not 200: %d", httpData.Status())
	}
	if res.Token != "xyz123" {
		t.Errorf("expected token xyz123, got %s", res.Token)
	}
}
