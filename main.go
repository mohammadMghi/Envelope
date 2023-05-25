package envlope

import (
 
	"net/http"
)

func main() {

	m := New()
      commonMiddleware := []Middleware{
            CurrentTimeHandler,
         }
      b := bar()
      m.r.POST("/",m.MultipleMiddleware(  b  , commonMiddleware   ))

      http.ListenAndServe(":8080" , m)
}
func logging() http.HandlerFunc {
      return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            
      })
  }

  func bar() http.HandlerFunc{
      return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            
      })
  }

  func CurrentTimeHandler(w http.ResponseWriter, r *http.Request)  {
      curTime := time.Now().Format(time.Kitchen)
      w.Write([]byte(fmt.Sprintf("the current time is %v", curTime)))
  }