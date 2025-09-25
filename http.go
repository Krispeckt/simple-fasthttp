package simple_fasthttp

type Http[E any] interface {
	Status() int
	Headers() map[string]string
	Raw() string
	Payload() *E
}

type HttpWrapper[E any] struct {
	status  int
	headers map[string]string
	payload *E
	raw     string
}

func (h *HttpWrapper[E]) Status() int {
	return h.status
}

func (h *HttpWrapper[E]) Headers() map[string]string {
	return h.headers
}

func (h *HttpWrapper[E]) Raw() string {
	return h.raw
}

func (h *HttpWrapper[E]) Payload() *E {
	return h.payload
}
