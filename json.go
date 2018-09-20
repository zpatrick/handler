package handler

import (
	"encoding/json"
	"net/http"
)

// JSON returns a http.Handler which writes the specified status and body in JSON format.
func JSON(status int, body interface{}) http.Handler {
	header := http.Header{}
	header.Set("Content-Type", "application/json")
	return JSONWithHeader(status, body, header)
}

// JSONWithHeader returns a http.Handler which writes the specified status, body, and header in JSON format.
func JSONWithHeader(status int, body interface{}, header http.Header) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		addHeader(w, header)
		w.WriteHeader(status)
		if err := json.NewEncoder(w).Encode(body); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
