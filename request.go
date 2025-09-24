package simple_fasthttp

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"
	"time"

	"github.com/valyala/fasthttp"
)

type RequestOptions struct {
	Method  string
	URL     *url.URL
	Params  url.Values
	Headers map[string]string
	Body    any
	Timeout time.Duration
}

func Do[T any, E any](ctx context.Context, opts RequestOptions) (*T, Error, error) {
	if opts.URL == nil {
		return nil, nil, errors.New("url is required")
	}
	if opts.Method == "" {
		opts.Method = fasthttp.MethodGet
	}
	if opts.Timeout == 0 {
		opts.Timeout = 5 * time.Second
	}

	if opts.Params != nil {
		opts.URL.RawQuery = opts.Params.Encode()
	}

	client := fasthttp.Client{}
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod(opts.Method)
	req.SetRequestURI(opts.URL.String())

	for k, v := range opts.Headers {
		req.Header.Set(k, v)
	}

	if opts.Body != nil {
		data, err := json.Marshal(opts.Body)
		if err != nil {
			return nil, nil, err
		}
		req.SetBody(data)
		req.Header.Set("Content-Type", "application/json")
	}

	// Выполняем с таймаутом
	if err := client.DoTimeout(req, resp, opts.Timeout); err != nil {
		// Проверяем отмену контекста
		select {
		case <-ctx.Done():
			return nil, nil, ctx.Err()
		default:
		}
		return nil, nil, err
	}

	status := resp.StatusCode()
	body := resp.Body()

	if status < 200 || status >= 300 {
		if len(body) == 0 {
			return nil, &ErrorWrapper[E]{status: status, raw: ""}, nil
		}
		var parsed E
		if err := json.Unmarshal(body, &parsed); err == nil {
			return nil, &ErrorWrapper[E]{status: status, payload: &parsed, raw: string(body)}, nil
		}
		return nil, &ErrorWrapper[E]{status: status, raw: string(body)}, nil
	}

	if len(body) == 0 {
		return new(T), nil, nil
	}

	var result T
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, nil, err
	}
	return &result, nil, nil
}
