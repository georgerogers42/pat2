# pat.go - A Sinatra style pattern muxer for Go's net/http library

## INSTALL

	$ go install github.com/bmizerany/pat.go

## USE

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

It's that simple.

## CONTRIBUTORS

* Keith Rarick (@krarick) - github.com/kr
* Blake Mizerany (@bmizerany) - github.com/bmizerany
* Evan Shaw
* George Rogers - github.com/georgerogers42
* Matt Nunogawa (@amattn) - github.com/amattn

## LICENSE

Copyright (C) 2012 by Keith Rarick, Blake Mizerany

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE. 
