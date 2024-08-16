package errors

import (
	"fmt"
	"net/http"
)

type NotFound struct {
	ID string
}

func (e NotFound) Error() string {
	return fmt.Sprintf("Entity not found. ID: '%s'", e.ID)
}

func (NotFound) StatusCode() int {
	return http.StatusBadRequest
}
