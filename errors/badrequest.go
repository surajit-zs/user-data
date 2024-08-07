package errors

import "net/http"

type BadRequest struct {
	Err error
}

func (e BadRequest) Error() string {
	return e.Err.Error()
}

func (BadRequest) StatusCode() int {
	return http.StatusBadRequest
}
