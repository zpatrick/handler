package handler

import "net/http"

// A Constructor constructs a http.Handler from a given *http.Request.
type Constructor func(r *http.Request) http.Handler

// ServeHTTP satisfies the http.Handler interface.
func (c Constructor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c(r).ServeHTTP(w, r)
}
