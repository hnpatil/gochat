package errors

import (
	"fmt"
	"github.com/hnpatil/gochat/entities/roommember"
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

func EntityExists(entity string) error {
	return response{
		message: fmt.Sprintf("%s already exists", entity),
		status:  http.StatusConflict,
		level:   logging.WARN,
	}
}

func Forbidden(operation string) error {
	return response{
		message: fmt.Sprintf("Cannot %s", operation),
		status:  http.StatusForbidden,
		level:   logging.WARN,
	}
}

func UnAuthorised(operation, resource string) error {
	return response{
		message: fmt.Sprintf("Cannot %s %s", operation, resource),
		status:  http.StatusUnauthorized,
		level:   logging.WARN,
	}
}

func MissingRoles(roles []roommember.Role) error {
	return response{
		message: fmt.Sprintf("User lacks one of the required roles - %s", roles),
		status:  http.StatusUnauthorized,
		level:   logging.WARN,
	}
}
