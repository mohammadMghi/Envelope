package envelope

import (
	"net/http"
 
)

type Params struct{

}


func NewParams() Params{
	return Params{}
}

 

func logger() http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	  print("logger test middleware")
	})
}