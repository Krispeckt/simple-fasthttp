package simple_fasthttp

import (
	"context"
	"net/url"
	"testing"
	"time"
)

func TestDoErrorNoURL(t *testing.T) {
	opts := RequestOptions{}
	_, _, err := Do[struct{}, Error](context.Background(), opts)
	if err == nil {
		t.Error("expected error when URL is nil")
	}
}

func TestDoErrorTimeout(t *testing.T) {
	u, _ := url.Parse("http://10.255.255.1") // unroutable
	opts := RequestOptions{
		URL:     u,
		Timeout: 1 * time.Millisecond,
	}

	_, _, err := Do[struct{}, Error](context.Background(), opts)
	if err == nil {
		t.Error("expected timeout error")
	}
}
