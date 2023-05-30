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

      r.GET("/asfdsf" , nil)
      return r
    })
  
      // envelop.Router.POST("/",b)

      
      
      
      http.ListenAndServe(envelop.Port , envelop)
}

 

 

  func test1(http.HandlerFunc) http.HandlerFunc{
      return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        print("asdasssssssssd")
      })
  }
 