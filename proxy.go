package simple_fasthttp

import (
	"fmt"
	"net/url"
	"time"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

type ProxyOptions struct {
	// URL types:
	//   http://host:port
	//   http://user:pass@host:port
	//   socks5://host:port
	//   socks5://user:pass@host:port
	URL *url.URL

	// HTTP_PROXY/HTTPS_PROXY/NO_PROXY
	UseEnv bool

	DualStack bool
}

func (o *ProxyOptions) NewClient(timeout time.Duration) (*fasthttp.Client, error) {
	client := fasthttp.Client{}
	switch {
	case o.UseEnv:
		client.Dial = fasthttpproxy.FasthttpProxyHTTPDialerTimeout(timeout)
	case o.URL != nil:
		switch o.URL.Scheme {
		case "socks5", "socks5h":
			if o.DualStack {
				client.Dial = fasthttpproxy.FasthttpSocksDialerDualStack(o.URL.String())
			} else {
				client.Dial = fasthttpproxy.FasthttpSocksDialer(o.URL.String())
			}
		case "http", "https":
			arg := httpProxyArg(o.URL)
			if o.DualStack {
				client.Dial = fasthttpproxy.FasthttpHTTPDialerDualStackTimeout(arg, timeout)
			} else {
				client.Dial = fasthttpproxy.FasthttpHTTPDialerTimeout(arg, timeout)
			}
		default:
			return nil, fmt.Errorf("unsupported proxy scheme: %q", o.URL.Scheme)
		}
	}
	return &client, nil
}

func httpProxyArg(u *url.URL) string {
	if u == nil {
		return ""
	}
	if u.User != nil {
		if p, ok := u.User.Password(); ok {
			return fmt.Sprintf("%s:%s@%s", u.User.Username(), p, u.Host)
		}
		return fmt.Sprintf("%s@%s", u.User.Username(), u.Host)
	}
	return u.Host
}
