package envlope

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)
 
type Log struct {
	hader map[string][]string `json:"hader"`
	wroteHeader bool `json:"wroteHeader"`
	time time.Time `json:"time"`
	realIp string `json:"realIp"`
	ip string `json:"ip"]`
}

func NewLog () Log{
	return Log{}
}

func  (log *Log)RequestLogger(http.HandlerFunc) http.HandlerFunc{
 
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeNow := time.Now()

		log.hader = w.Header().Clone()
		
		log.time = timeNow

		log.realIp = r.Header.Get("X-Real-Ip")

		log.ip = r.Header.Get("X-Forwarded-For")

		myLog, _ := json.Marshal(log)
		fmt.Println(myLog)
		


	})
}