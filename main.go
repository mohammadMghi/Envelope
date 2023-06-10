package main

import (
	"fmt"
	"net/http"

	"github.com/mohammadmghi/envelope/envelope"
)

type test struct{
  ItemOne string
  itemTwo string
}

func main() {
 
	envelop := envelope.New(":8081")
 


  cache :=  envelope.NewCacheEnv(100000000000000)

  myTest :=test{
    "Hello" , 
    "RR",
  }
  
 
    envelop.Router.POST("/" , func() string{

      return ""
    })
  
  
    envelop.Router.POST("/getTest" , func() string{
      cache.Set("Test" , myTest ,100000000000000 )
      return ""
    })






    envelop.RouterGroup.Group("/test" , func(r envelope.RouterGroup) envelope.RouterGroup {
    

      r.POST("/sdf" ,  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        inter , b  :=cache.Get("Test")
            fmt.Printf("this is iiiiiiiiiiiiiiiiii %+v\n",inter )
        if b {
          fmt.Printf("this is iiiiiiiiiiiiiiiiii %+v\n",inter )
        }
        
  
      }))

      r.POST("/asfdsf" , func() string{
 
        cache.Set("Test" , myTest ,100000000000000 )
     
        return "sdasdas"

      })

 
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
