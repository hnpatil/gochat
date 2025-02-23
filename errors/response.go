package errors

import (
	"fmt"
	"gofr.dev/pkg/gofr/logging"
	"net/http"
)

type response struct {
	message string
	status  int
	level   logging.Level
}

func (e response) Error() string {
	return e.message
}

func (e response) StatusCode() int {
	return e.status
}

func (e response) LogLevel() logging.Level {
	return e.level
}

func EntityNotFound(entity string) error {
	return response{
		message: fmt.Sprintf("%s not found", entity),
		status:  http.StatusNotFound,
		level:   logging.WARN,
	}
}

func UnAuthorised(operation, resource string) error {
	return response{
		message: fmt.Sprintf("User cannot %s on %s", operation, resource),
		status:  http.StatusUnauthorized,
		level:   logging.WARN,
	}
}
