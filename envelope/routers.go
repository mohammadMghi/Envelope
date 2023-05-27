package envelope

import (
 
 
	"net/http"
	"regexp"
 
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
	PathRootGroup PathRootGroup
}

type PathGroup struct {
    leftPath  *PathGroup
    rightPath *PathGroup
	Root string
	Path string
	Method string
	Handler http.Handler
}
 
type PathRootGroup struct {
    root *PathGroup
}

 

func (router *Router) Group(path string  ,fn func(r Router) Router ) {

		router.PathRootGroup.root = &PathGroup{leftPath: nil , rightPath: nil , Path : path}


		for _, route := range fn(*router).routes{
			if router.PathRootGroup.root.rightPath != nil{
				router.PathRootGroup.root.rightPath = &PathGroup{Path : route.Pattern, rightPath: nil , leftPath : nil }
			}else{
				router.PathRootGroup.root.leftPath = &PathGroup{Path : route.Pattern, rightPath: nil , leftPath : nil  }
			}

		}	
}



func (r *Router) InsertRootPath(path string  ,routerCallBack func(r Router) Router ){
	if path == pathRoot {
		 
	}
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

 



func (r *Router) DELETE(path string, handler http.HandlerFunc) {
	r.AddRoute(MethodDelete, path, handler)
}

func (r *Router) GET(path string, handler http.HandlerFunc) {
	r.AddRoute(MethodGet, path, handler)
}

func (r *Router) POST(path string, handler http.HandlerFunc) {
	r.AddRoute(MethodPost, path, handler)
}

func (r *Router) PUT(path string, handler http.HandlerFunc) {
	r.AddRoute(MethodPut, path, handler)
}

 

 

 


