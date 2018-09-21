package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zpatrick/handler"
)

func Index(r *http.Request) http.Handler {
	return handler.String(http.StatusOK, "Hello, World!")
}

func main() {
	r := mux.NewRouter()
	r.Handle("/", handler.Constructor(Index))

	log.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
