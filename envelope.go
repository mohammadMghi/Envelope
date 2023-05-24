package envlope

import (
	"log"
	"net/http"
	"reflect"
	"time"
)


type Envlope struct{
	handler	http.Handler
 	eHanlders []EHandler
	r Router
 
}

type EHandler interface{}

func New() *Envlope{
	router := NewRouter()
	return &Envlope{
		 r: *router,
	}
}


func (e *Envlope)addHandler(handler EHandler){
	// Checks if middlware handler is a function
	if reflect.TypeOf(handler).Kind() != reflect.Func{
		panic("type must be a callable function")
	}

	e.eHanlders = append(e.eHanlders , handler )
}



func (e *Envlope)addHandlers(handler ...EHandler){

	for _ , handler := range(handler){
		e.addHandler(handler)
	}

 
}

func (l *Envlope) ServeHTTP(w http.ResponseWriter , r *http.Request){
 
    l.handler.ServeHTTP(w, r)
}