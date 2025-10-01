package simple_fasthttp

import (
	"context"
	"net/url"
	"testing"
	"time"
)

type WithProxyTestResp struct {
	Args struct {
		Ping string `json:"ping"`
	} `json:"args"`
	Headers struct {
		Host            string `json:"host"`
		XRequestStart   string `json:"x-request-start"`
		Connection      string `json:"connection"`
		XForwardedProto string `json:"x-forwarded-proto"`
		XForwardedPort  string `json:"x-forwarded-port"`
		XAmznTraceId    string `json:"x-amzn-trace-id"`
		SecFetchDest    string `json:"sec-fetch-dest"`
		UserAgent       string `json:"user-agent"`
		Accept          string `json:"accept"`
		SecFetchSite    string `json:"sec-fetch-site"`
		SecFetchMode    string `json:"sec-fetch-mode"`
		AcceptLanguage  string `json:"accept-language"`
		Priority        string `json:"priority"`
		AcceptEncoding  string `json:"accept-encoding"`
	} `json:"headers"`
	Url string `json:"url"`
}

func TestRequestWithProxy(t *testing.T) {
	u, _ := url.Parse("https://postman-echo.com/get?ping=pong")
	opts := RequestOptions{
		URL:     u,
		Timeout: 60 * time.Second,
		Proxy:   &ProxyOptions{UseEnv: true},
	}

	res, httpData, err := Do[WithProxyTestResp, Http[any]](context.Background(), opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if httpData.Status() != 200 {
		t.Fatalf("unexpected http error: %v", httpData.Payload())
	}
	if res.Args.Ping != "pong" {
		t.Fatalf("unexpected ping value: %v", res.Args.Ping)
	}
}
