package main

import (
    "github.com/mohammadmghi/envelope/envelope"
	"net/http"
)

func main() {
 
	envelop := envelope.New(":8080")
    logger := envelope.NewLog()
    b := logger.RequestLogger(func(w http.ResponseWriter, r *http.Request) {
      
    })
  

  
      envelop.Router.POST("/",b)

      
      
      
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
 