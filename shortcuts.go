package simple_fasthttp

import (
	"context"
	"net/url"

	"github.com/valyala/fasthttp"
)

//
// ====== GET ======
//

func GetWithParseErr[T any, E any](ctx context.Context, u *url.URL, params url.Values, headers map[string]string) (*T, Http[E], error) {
	return Do[T, E](ctx, RequestOptions{
		Method:  fasthttp.MethodGet,
		URL:     u,
		Params:  params,
		Headers: headers,
	})
}

func Get[T any](ctx context.Context, u *url.URL, params url.Values, headers map[string]string) (*T, Http[any], error) {
	return Do[T, any](ctx, RequestOptions{
		Method:  fasthttp.MethodGet,
		URL:     u,
		Params:  params,
		Headers: headers,
	})
}

func GetWithParams[T any](ctx context.Context, u *url.URL, params url.Values) (*T, Http[any], error) {
	return Do[T, any](ctx, RequestOptions{
		Method: fasthttp.MethodGet,
		URL:    u,
		Params: params,
	})
}

func GetWithHeaders[T any](ctx context.Context, u *url.URL, headers map[string]string) (*T, Http[any], error) {
	return Do[T, any](ctx, RequestOptions{
		Method:  fasthttp.MethodGet,
		URL:     u,
		Headers: headers,
	})
}

func GetSimple[T any](ctx context.Context, u *url.URL) (*T, Http[any], error) {
	return Do[T, any](ctx, RequestOptions{
		Method: fasthttp.MethodGet,
		URL:    u,
	})
}

//
// ====== POST ======
//

func PostWithParseErr[T any, E any](ctx context.Context, u *url.URL, body any, params url.Values, headers map[string]string) (*T, Http[E], error) {
	return Do[T, E](ctx, RequestOptions{
		Method:  fasthttp.MethodPost,
		URL:     u,
		Body:    body,
		Params:  params,
		Headers: headers,
	})
}

func Post[T any](ctx context.Context, u *url.URL, body any, params url.Values, headers map[string]string) (*T, Http[any], error) {
	return Do[T, any](ctx, RequestOptions{
		Method:  fasthttp.MethodPost,
		URL:     u,
		Body:    body,
		Params:  params,
		Headers: headers,
	})
}

func PostWithParams[T any](ctx context.Context, u *url.URL, body any, params url.Values) (*T, Http[any], error) {
	return Do[T, any](ctx, RequestOptions{
		Method: fasthttp.MethodPost,
		URL:    u,
		Body:   body,
		Params: params,
	})
}

func PostWithBody[T any](ctx context.Context, u *url.URL, body any) (*T, Http[any], error) {
	return Do[T, any](ctx, RequestOptions{
		Method: fasthttp.MethodPost,
		URL:    u,
		Body:   body,
	})
}

//
// ====== PUT ======
//

func PutWithParseErr[T any, E any](ctx context.Context, u *url.URL, body any, headers map[string]string) (*T, Http[E], error) {
	return Do[T, E](ctx, RequestOptions{
		Method:  fasthttp.MethodPut,
		URL:     u,
		Body:    body,
		Headers: headers,
	})
}

func Put[T any](ctx context.Context, u *url.URL, body any, headers map[string]string) (*T, Http[any], error) {
	return Do[T, any](ctx, RequestOptions{
		Method:  fasthttp.MethodPut,
		URL:     u,
		Body:    body,
		Headers: headers,
	})
}

func PutWithBody[T any](ctx context.Context, u *url.URL, body any) (*T, Http[any], error) {
	return Do[T, any](ctx, RequestOptions{
		Method: fasthttp.MethodPut,
		URL:    u,
		Body:   body,
	})
}

//
// ====== DELETE ======
//

func DeleteWithParseErr[T any, E any](ctx context.Context, u *url.URL, headers map[string]string) (*T, Http[E], error) {
	return Do[T, E](ctx, RequestOptions{
		Method:  fasthttp.MethodDelete,
		URL:     u,
		Headers: headers,
	})
}

func Delete[T any](ctx context.Context, u *url.URL, headers map[string]string) (*T, Http[any], error) {
	return Do[T, any](ctx, RequestOptions{
		Method:  fasthttp.MethodDelete,
		URL:     u,
		Headers: headers,
	})
}

func DeleteNoHeaders[T any](ctx context.Context, u *url.URL) (*T, Http[any], error) {
	return Do[T, any](ctx, RequestOptions{
		Method: fasthttp.MethodDelete,
		URL:    u,
	})
}
