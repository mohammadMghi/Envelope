package envelope

import (
	"fmt"
	"net/http"
	"time"
)
 
type Log struct {
	Hader map[string][]string `json:"hader"`
	WroteHeader bool `json:"wroteHeader"`
	Time time.Time `json:"time"`
	RealIp string `json:"realIp"`
	Ip string `json:"ip"`
}

func NewLog () Log{
	return Log{}
}

func  (log *Log)RequestLogger(  http.HandlerFunc) http.HandlerFunc{
 
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
 
		
		log.Hader = w.Header().Clone()
 
		log.Time = time.Now()
		log.RealIp = r.Header.Get("X-Real-Ip")

		log.Ip = r.Header.Get("X-Forwarded-For")

		
		fmt.Printf("%+v\n", log)
		
		

	})
}