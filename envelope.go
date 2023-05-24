package envlope

import (
 
	"reflect"
)


type Envlope struct{
 	eHanlders []EHandler
 
}

type EHandler struct{}

func New() *Envlope{
 
	return &Envlope{
		 
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