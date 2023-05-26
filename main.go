package envlope

import (
 
	"net/http"
)

func main() {

	m := New()
      commonMiddleware := []Middleware{
             test,
             test1,
         }
      b := test(func(w http.ResponseWriter, r *http.Request) {
            
      })
      m.r.POST("/",m.MultipleMiddleware(  b  ,   commonMiddleware  ))

      http.ListenAndServe(":8080" , m)
}

 

func test(http.HandlerFunc) http.HandlerFunc{
      return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            
      })
  }

  func test1(http.HandlerFunc) http.HandlerFunc{
      return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            
      })
  }
 