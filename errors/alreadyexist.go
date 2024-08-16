package errors

import (
	"fmt"
	"net/http"
)

type AlreadyExist struct {
	UserName string
}

func (a AlreadyExist) Error() string {
	return fmt.Sprintf("Entity already exists user name '%s'", a.UserName)
}

func (AlreadyExist) StatusCode() int {
	return http.StatusConflict
}
