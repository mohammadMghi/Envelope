package envlope

import (
 
	"net/http"
)

func main() {

	r := NewRouter()
      r.GET("",bazHandler()) 
  


      http.ListenAndServe(":8080" , r)
}


func bazHandler() http.Handler {
      return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
 
      })
  }