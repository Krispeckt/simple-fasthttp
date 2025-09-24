package simple_fasthttp

import (
	"context"
	"net/url"

	"github.com/valyala/fasthttp"
)

func Get[T any](ctx context.Context, u *url.URL, params url.Values, headers map[string]string) (*T, Error, error) {
	return Do[T, Error](ctx, RequestOptions{
		Method:  fasthttp.MethodGet,
		URL:     u,
		Params:  params,
		Headers: headers,
	})
}

func GetWithParams[T any](ctx context.Context, u *url.URL, params url.Values) (*T, Error, error) {
	return Do[T, Error](ctx, RequestOptions{
		Method: fasthttp.MethodGet,
		URL:    u,
		Params: params,
	})
}

func GetWithHeaders[T any](ctx context.Context, u *url.URL, headers map[string]string) (*T, Error, error) {
	return Do[T, Error](ctx, RequestOptions{
		Method:  fasthttp.MethodGet,
		URL:     u,
		Headers: headers,
	})
}

func GetSimple[T any](ctx context.Context, u *url.URL) (*T, Error, error) {
	return Do[T, Error](ctx, RequestOptions{
		Method: fasthttp.MethodGet,
		URL:    u,
	})
}

func Post[T any](ctx context.Context, u *url.URL, body any, params url.Values, headers map[string]string) (*T, Error, error) {
	return Do[T, Error](ctx, RequestOptions{
		Method:  fasthttp.MethodPost,
		URL:     u,
		Body:    body,
		Params:  params,
		Headers: headers,
	})
}

func PostWithParams[T any](ctx context.Context, u *url.URL, body any, params url.Values) (*T, Error, error) {
	return Do[T, Error](ctx, RequestOptions{
		Method: fasthttp.MethodPost,
		URL:    u,
		Body:   body,
		Params: params,
	})
}

func PostWithBody[T any](ctx context.Context, u *url.URL, body any) (*T, Error, error) {
	return Do[T, Error](ctx, RequestOptions{
		Method: fasthttp.MethodPost,
		URL:    u,
		Body:   body,
	})
}

func Put[T any](ctx context.Context, u *url.URL, body any, headers map[string]string) (*T, Error, error) {
	return Do[T, Error](ctx, RequestOptions{
		Method:  fasthttp.MethodPut,
		URL:     u,
		Body:    body,
		Headers: headers,
	})
}

func PutWithBody[T any](ctx context.Context, u *url.URL, body any) (*T, Error, error) {
	return Do[T, Error](ctx, RequestOptions{
		Method: fasthttp.MethodPut,
		URL:    u,
		Body:   body,
	})
}

func Delete[T any](ctx context.Context, u *url.URL, headers map[string]string) (*T, Error, error) {
	return Do[T, Error](ctx, RequestOptions{
		Method:  fasthttp.MethodDelete,
		URL:     u,
		Headers: headers,
	})
}

func DeleteNoHeaders[T any](ctx context.Context, u *url.URL) (*T, Error, error) {
	return Do[T, Error](ctx, RequestOptions{
		Method: fasthttp.MethodDelete,
		URL:    u,
	})
}
