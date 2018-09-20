package handler

import (
	"net/http"
)

// Error returns a http.Handler which writes the specified status and error.
func Error(status int, err error) http.Handler {
	return String(status, err.Error())
}

// Error returns a http.Handler which writes the specified status and formatted error.
func Errorf(status int, format string, tokens ...interface{}) http.Handler {
	return Stringf(status, format, tokens...)
}

// ErrorWithHeader returns a http.Handler which writes the specified status, error, and header.
func ErrorWithHeader(status int, err error, header http.Header) http.Handler {
	return StringWithHeader(status, err.Error(), header)
}
