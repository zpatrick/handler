package handler

import (
	"net/http"
)

// Bytes returns a http.Handler which writes the specified status and body.
func Bytes(status int, body []byte) http.Handler {
	return BytesWithHeader(status, body, http.Header{})
}

// BytesWithHeader returns a http.Handler which writes the specified status, body, and header.
func BytesWithHeader(status int, body []byte, header http.Header) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		addHeader(w, header)
		w.WriteHeader(status)
		if _, err := w.Write(body); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
