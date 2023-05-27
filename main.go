package main

import (
    "github.com/mohammadmghi/envelope/envlope"
	"net/http"
)

func main() {
 
	envelop := envlope.New()
      commonMiddleware := []envlope.Middleware{
             test,
             test1,
         }

    //   logger := NewLog()
      b := test(func(w http.ResponseWriter, r *http.Request) {
            
      })
      envelop.Router.POST("/",envelop.MultipleMiddleware(  b,   commonMiddleware  ))


      
      http.ListenAndServe(":8080" , envelop)
}

 

func test(http.HandlerFunc) http.HandlerFunc{
      return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            
      })
  }

  func test1(http.HandlerFunc) http.HandlerFunc{
      return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            
      })
  }
 