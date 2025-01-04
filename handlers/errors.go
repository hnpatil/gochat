package handlers

import (
	"gofr.dev/pkg/gofr/logging"
	"net/http"
)

type NotImplemented struct {
	error string
}

func (n NotImplemented) Error() string {
	return "Not Implemented"
}

func (n NotImplemented) StatusCode() int {
	return http.StatusNotImplemented
}

func (n NotImplemented) LogLevel() logging.Level {
	return logging.WARN
}
