// pat.go a sinatra like muxer for go.
package pat

import (
	"net/http"
)

type Handler interface {
	Handle(params Params, splat string) http.HandlerFunc
}
type HandlerFunc func (params Params, splat string) http.HandlerFunc
func (h HandlerFunc) Handle(params Params, splat string) http.HandlerFunc {
	return h(params, splat)
}

type Params map[string]string

type PatternServeMux struct {
	handlers map[string][]*patHandler
}

// Creates an new *PatternServeMux.
func New() *PatternServeMux {
	return &PatternServeMux{make(map[string][]*patHandler)}
}

// Implements HttpHandler.
func (p *PatternServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, ph := range p.handlers[r.Method] {
		if params, splat, ok := ph.try(r.URL.Path); ok {
			px := ph.Handle(params, splat)
			px.ServeHTTP(w, r)
			return
		}
	}

	http.NotFound(w, r)
}

// Adds the pattern handler pair for a Get.
func (p *PatternServeMux) Get(pat string, h Handler) {
	p.Add("GET", pat, h)
}

// Adds the pattern handler pair for a Post.
func (p *PatternServeMux) Post(pat string, h Handler) {
	p.Add("POST", pat, h)
}

// Adds the pattern handler pair for a Put.
func (p *PatternServeMux) Put(pat string, h Handler) {
	p.Add("PUT", pat, h)
}

// Adds the pattern handler pair for a Delete.
func (p *PatternServeMux) Del(pat string, h Handler) {
	p.Add("DELETE", pat, h)
}

// Adds the pattern handler pair for a HTTP Method meth.
func (p *PatternServeMux) Add(meth, pat string, h Handler) {
	p.handlers[meth] = append(p.handlers[meth], &patHandler{pat, h})
}

type patHandler struct {
	pat string
	Handler
}

func (ph *patHandler) try(path string) (Params, string, bool) {
	p := make(Params)
	var i, j int
	for i < len(path) {
		switch {
		case j == len(ph.pat) && ph.pat[j-1] == '/':
			// Should i put a special form variable splat for this case
			return p, path[i:], true
		case j >= len(ph.pat):
			return nil, "",false
		case ph.pat[j] == ':':
			var name, val string
			name, j = find(ph.pat, '/', j)
			val, i = find(path, '/', i)
			p[name] = val
		case path[i] == ph.pat[j]:
			i++
			j++
		default:
			return nil, "", false
		}
	}
	if j != len(ph.pat) {
		return nil, "", false
	}
	return p, "", true
}

func find(s string, c byte, i int) (string, int) {
	j := i
	for j < len(s) && s[j] != c {
		j++
	}
	return s[i:j], j
}
