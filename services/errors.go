package services

import (
	"gofr.dev/pkg/gofr/logging"
	"net/http"
)

type Error struct {
	error  string
	status int
	level  logging.Level
}

func (e Error) Error() string {
	return e.error
}

func (e Error) StatusCode() int {
	return e.status
}

func (e Error) LogLevel() logging.Level {
	return e.level
}

func UnAuthorisedError(message string) Error {
	return Error{
		error:  message,
		status: http.StatusUnauthorized,
		level:  logging.WARN,
	}
}
