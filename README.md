# Handler

[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/zpatrick/handler/blob/master/LICENSE)
[![Go Doc](https://godoc.org/github.com/zpatrick/handler?status.svg)](https://godoc.org/github.com/zpatrick/handler)

## Usage
Handler is a collection of common `http.Handlers` used in Go web applications, typically REST APIs. 
These handlers integrate well with applications that use the [Handler Constructor](https://godoc.org/github.com/zpatrick/handler#Constructor) pattern: 

```go
type Constructor func(r *http.Request) http.Handler
```

Please see the [example](/example) for a working example. 

### Why use the Handler Constructor pattern? 
Well, it depends. Take a look at the following examples: 

#### Using the conventional pattern
```go
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p *Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	if err := db.Insert(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
```

#### Using the Handler Constructor pattern
```go
func CreateProduct(r *http.Request) http.Handler {
	var p *Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		return handler.Error(http.StatusBadRequest, err)
	}
	
	if err := db.Insert(p); err != nil {
		return handler.Error(http.StatusInternalServerError, err)
	}
	
	return handler.JSON(http.StatusCreated, p)
}
```

Hopefully, at least two things stick out:
* The Handler Constructor pattern allows for early function departures without requiring an empty `return` statement.
This should feel intuitive with how much we encounter `return err` statements. 
* We don't even see `w http.ResponseWriter` in our constructor, making it much more difficult to make the mistake of writing to `w` when we shouldn't (e.g. multiple times). 
It is easier for this type of mistake to occur using the conventional pattern, say, forgetting an empty `return` statement after calling `http.Error`. Note that you _can_ get access to `w` in your constructor by returning a [http.HandlerFunc](https://golang.org/pkg/net/http/#HandlerFunc). 

#### When to not use the Handler Constructor pattern
The points made above are enough to convince me to use this pattern for most of my Go applications; but not all of them.
Here are some reasons you may want avoid using the handler.Constructor pattern: 
* The size/scope of your project is small. 
* You need to stream data to the `http.ResponseWriter`. 
* You find yourself frequently needing access to `w http.ResponseWriter` directly - to the point where the majority of your constructors become nothing more than wrappers around [http.HandlerFuncs](https://golang.org/pkg/net/http/#HandlerFunc). 


### Routing 
The [Constructor](https://godoc.org/github.com/zpatrick/handler#Constructor) type satisfies the [http.Handler](https://golang.org/pkg/net/http/#Handler) interface,
so they can be used anywhere a `http.Handler` can. 
Here is an example using Handler Constructors with [gorilla/mux](https://github.com/gorilla/mux):

```go
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
```
