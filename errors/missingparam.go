package errors

import (
	"net/http"
	"strings"
)

type MissingParam struct {
	Params []string
}

func (e MissingParam) Error() string {
	switch len(e.Params) {
	case 0:
		return "missing param"
	case 1:
		return "missing param: " + e.Params[0]
	default:
		return "missing params: " + strings.Join(e.Params, ", ")
	}
}

func (MissingParam) StatusCode() int {
	return http.StatusBadRequest
}
