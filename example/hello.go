package main

import (
	"github.com/georgerogers42/pat2"
	"io"
	"net/http"
)

func main() {
	m := pat.New()
	m.Get("/hello/:name", pat.HandlerFunc(hello))
	m.Get("/splat/", pat.HandlerFunc(splat))
	http.ListenAndServe("localhost:5000", m)
}

func hello(params pat.Params, _ string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Path variable names are in the URL.Query() and start with ':'.
		name := params[":name"]
		io.WriteString(w, "Hello, "+name)
	}
}

func splat(_ pat.Params, s string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Path variable names are in the URL.Query() and start with ':'.
		io.WriteString(w, "Splat: "+s)
	}
}
