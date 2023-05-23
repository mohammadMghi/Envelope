package envlope

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

 
const (
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH" // RFC 5789
	MethodDelete  = "DELETE"
	MethodConnect = "CONNECT"
	MethodOptions = "OPTIONS"
	MethodTrace   = "TRACE"
	pathRoot      string = "/"
	pathDelimiter string = "/"
)
 

 
func NewRouter() *Router {
	return &Router{}
}

type Route struct {
	Method  string
	Pattern string
	Handler http.Handler
}

type Router struct {
	routes []Route
}



func (r *Router) AddRoute(method, path string, handler http.Handler) {
	r.routes = append(r.routes, Route{Method: method, Pattern: path, Handler: handler})
}

func (r *Router) getHandler(method, path string) http.Handler {
	for _, route := range r.routes {
		re := regexp.MustCompile(route.Pattern)
		if route.Method == method && re.MatchString(path) {
			return route.Handler
		}
	}
	return http.NotFoundHandler()
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method

	handler := r.getHandler(method, path)

	// handler middlewares go here

	handler.ServeHTTP(w, req)
}

func (r *Router) DELETE(path string, handler http.Handler) {
	r.AddRoute("DELETE", path, handler)
}

func (r *Router) GET(path string, handler http.Handler) {
	r.AddRoute("GET", path, handler)
}

func (r *Router) POST(path string, handler http.Handler) {
	r.AddRoute("POST", path, handler)
}

func (r *Router) PUT(path string, handler http.Handler) {
	r.AddRoute("PUT", path, handler)
}

 

 

