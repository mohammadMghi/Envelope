package envlope

import (
 
 
	"sync"

	"net/http"
	"reflect"
)



type Envlope struct{
	c sync.Pool
 	eHanlders []http.Handler
	Router Router
	log Log
 
}

type EHandler interface{}

func New( ) *Envlope{
	router := NewRouter()
	log := NewLog()
	return &Envlope{
		Router: *router,
		log: log,
	}
}


func (e *Envlope)addHandler(handler http.Handler){
	// Checks if middlware handler is a function
	if reflect.TypeOf(handler).Kind() != reflect.Func{
		panic("type must be a callable function")
	}

	e.eHanlders = append(e.eHanlders , handler )
}



func (e *Envlope)addHandlers(handler ...http.Handler){

	for _ , handler := range(handler){
		e.addHandler(handler)
	}

 
}

type Middleware func(http.HandlerFunc) http.HandlerFunc
 
 

func (l *Envlope) ServeHTTP(w http.ResponseWriter , req *http.Request){

	path := req.URL.Path
	method := req.Method

	handler := l.Router.getHandler(method, path)
 
	// handler middlewares go here

	handler.ServeHTTP(w, req)
 
}

