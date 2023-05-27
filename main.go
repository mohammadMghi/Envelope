package main

import (
    "github.com/mohammadmghi/envelope/envlope"
	"net/http"
)

func main() {
 
	envelop := envlope.New()
    logger := envlope.NewLog()
    b := logger.RequestLogger(func(w http.ResponseWriter, r *http.Request) {
      
    })
  

  
      envelop.Router.POST("/",envelop.MultipleMiddleware(  b , nil ))

      
      
      http.ListenAndServe(":8081" , envelop)
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
 