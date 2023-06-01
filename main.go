package main

import (
 
	"net/http"

	"github.com/mohammadmghi/envelope/envelope"
)

func main() {
 
	envelop := envelope.New(":8081")
 
 
  
    envelop.Router.Group("/test" , func(r envelope.Router) envelope.Router {
    
     
      r.POST("/sdf" ,  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        print("asdasd")
      }))

      r.POST("/asfdsf" , envelope.Chain(logger(), Method("POST") ))
      return r
    })
  
      // envelop.Router.POST("/",b)

      
      
      
      http.ListenAndServe(envelop.Port , envelop)
}

 

 

  func logger() http.HandlerFunc{
      return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        print("logger test middleware")
      })
  }
 
  func Method(m string) envelope.Middleware {
 
    return func(f http.HandlerFunc) http.HandlerFunc {
      print("Method Function is Post")
       return f
    }
}
