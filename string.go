package handler

import (
	"fmt"
	"net/http"
)

// String returns a http.Handler which writes the specified status and body.
func String(status int, body string) http.Handler {
	return BytesWithHeader(status, []byte(body), http.Header{})
}

// Stringf returns a http.Handler which writes the specified status and formatted string.
func Stringf(status int, format string, tokens ...interface{}) http.Handler {
	return String(status, fmt.Sprintf(format, tokens...))
}

// StringWithHeader returns a http.Handler which writes the specified status, body, and header.
func StringWithHeader(status int, body string, header http.Header) http.Handler {
	return BytesWithHeader(status, []byte(body), header)
}
