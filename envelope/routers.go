package envelope

import (
	"fmt"
	"net/http"

	"regexp"
	"strings"
	"unicode/utf8"

 
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

	Handler Handler
 
}
 

type Router struct {
	routes []Route
	PathGroup PathGroup
}

type PathGroup struct {
    leftPath  *PathGroup
    rightPath *PathGroup
	root *PathGroup
	Root string
	Path string
	Method string
	Handler Handler
	
}

type context struct{
	responseWriter http.ResponseWriter

}
 
 

func (router *Router) Group(path string  ,fn func(r Router) Router ) {
 
		router.PathGroup.root = &PathGroup{leftPath: nil , rightPath: nil , Path : path}


		for _, route := range fn(*router).routes{

			 
			if router.PathGroup.leftPath == nil{

			
				router.PathGroup.leftPath = &PathGroup{Path : route.Pattern,Handler: route.Handler , rightPath: nil , leftPath : nil  }

			}else{

			
				router.PathGroup.rightPath = &PathGroup{Path : route.Pattern, Handler: route.Handler ,rightPath: nil, leftPath : nil }

			}

		}
	
}
 
 

  
func (pathGroup *PathGroup) GetPathGroup(path string  ) *PathGroup {


	fmt.Printf("path group :: %+v\n", pathGroup.rightPath.Path)

	if pathGroup.leftPath.Path == "" {

		return nil
	}

 
	if pathGroup.rightPath.Path == "" {
 
		return nil
	}

	if  pathGroup.leftPath.Path ==  GetGroupPath(path){
 
		return pathGroup.leftPath
	}

	if pathGroup.rightPath.Path == GetGroupPath(path){
 
		return pathGroup.rightPath
	}



 
	if pathGroup.leftPath.Path == GetGroupPath(path){
 
		return pathGroup.leftPath.GetPathGroup(path)
	}else{
	 
		return pathGroup.rightPath.GetPathGroup(path)
	}
 
}

// func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
//     for _, m := range middlewares {
//         f = m(f)
//     }
//     return f
// }


func (r *Router) AddRoute(method, path string, handler Handler) {
	r.routes = append(r.routes, Route{Method: method, Pattern: path, Handler: handler})
}



func (r *Router) checkPathIsGroup (path string) bool{


	if r.PathGroup.root.Path == "/" + path{

		return true
	}

	return false
}

func (r *Router)getNormalHandler(path string, method string) Handler{

 
	for _, route := range r.routes {
				re := regexp.MustCompile(route.Pattern)
				if route.Method == method && re.MatchString(path) {
	
					return route.Handler
				}
			}

	return http.NotFoundHandler()
}


 
func (r *Router)getHandlerGroup(path string, method string) Handler{
	pathGroup := r.PathGroup.GetPathGroup(path)
 
	if  pathGroup.Handler != nil {
 
		return pathGroup.Handler
	}

	return http.NotFoundHandler()
}

 

func (r *Router) getHandler(method string, path string) Handler {
	fmt.Printf("path :: %+v\n", path)
	if path != pathRoot{
		isGroupPath := r.checkPathIsGroup(GetRootGroupPath(path))


		
		// Checks if root path exsited in group then get Handler
		if isGroupPath {

		pathGroup  := r.PathGroup.GetPathGroup(path)
		emptyPathGroup := PathGroup{}
			if  pathGroup !=&emptyPathGroup {
	
				return r.getHandlerGroup(path, method)
			
			}
		}
}
		
 
	return r.getNormalHandler(path , method)

}




func (r *Router) DELETE(path string, handler Handler) {
	r.AddRoute(MethodDelete, path, handler)
}

func (r *Router) GET(path string, handler Handler) {
	r.AddRoute(MethodGet, path, handler)
}

func (r *Router) POST(path string, handler Handler) {
	r.AddRoute(MethodPost, path, handler)
}

func (r *Router) PUT(path string, handler Handler) {
	r.AddRoute(MethodPut, path, handler)
}

 
func trimFirstRune(s string) string {
    _, i := utf8.DecodeRuneInString(s)
    return s[i:]
}
 

func  GetRootGroupPath(m string) string{
	m = trimFirstRune(m)
 
	if idx := strings.IndexByte(m, '/'); idx >= 0 {
		s := m[:idx]
	
		fmt.Println(s)
		return s
	} else {
 
		fmt.Println("Invalid string")
	}
	return ""
} 

func  GetGroupPath(m string) string{
	m = trimFirstRune(m)
 
	if idx := strings.IndexByte(m, '/'); idx >= 0 {
		s := m[idx:]
	
		fmt.Println(s)
		return s
	} else {
		fmt.Println("Invalid string")
	}
	return ""
} 

