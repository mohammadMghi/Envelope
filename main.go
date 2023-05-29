package main

import (
    "github.com/mohammadmghi/envelope/envelope"
	"net/http"
)

func main() {
 
	envelop := envelope.New(":8081")
    // logger := envelope.NewLog()
    // b := logger.RequestLogger(func(w http.ResponseWriter, r *http.Request) {
      
    // })
  
    envelop.Router.Group("/test" , func(r envelope.Router) envelope.Router {
      r.POST("/sdf" , nil)

      r.GET("/asfdsf" , nil)
      return r
    })
  
      // envelop.Router.POST("/",b)

      
      
      
      http.ListenAndServe(envelop.Port , envelop)
}

 

func test(http.HandlerFunc) http.HandlerFunc{
      return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            print("asdasd")
      })
  }

  func test1(http.HandlerFunc) http.HandlerFunc{
      return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        print("asdasssssssssd")
      })
  }
 