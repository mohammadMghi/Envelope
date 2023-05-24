package envlope

import (
 
	"net/http"
)

func main() {

	m := New()
    
  
      m.addHandlers(logging(),bar())

      http.ListenAndServe(":8080" , m)
}
func logging() http.HandlerFunc {
      return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            
      })
  }

  func bar() http.Handler{
      return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            
      })
  }

func bazHandler() http.Handler {
      return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            
      })
  }