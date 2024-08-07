package errors

import "net/http"

type DB struct {
	Err error
}

func (e DB) Error() string {
	return e.Err.Error()
}

func (DB) StatusCode() int {
	return http.StatusInternalServerError
}
