package simple_fasthttp

import (
	"encoding/json"
	"fmt"
)

type Error interface {
	error
	StatusCode() int
	RawBody() string
}

type ErrorWrapper[E any] struct {
	status  int
	payload *E
	raw     string
}

func (e *ErrorWrapper[E]) Error() string {
	if e.payload != nil {
		data, _ := json.Marshal(e.payload)
		return fmt.Sprintf("error (%d): %s", e.status, string(data))
	}
	return fmt.Sprintf("error (%d): %s", e.status, e.raw)
}

func (e *ErrorWrapper[E]) StatusCode() int { return e.status }
func (e *ErrorWrapper[E]) RawBody() string { return e.raw }
