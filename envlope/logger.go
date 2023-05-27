package envlope

import (
	"net/http"
	"time"
)
 
type Log struct {
	hader map[string][]string
	wroteHeader bool
	time time.Time
	realIp string
	ip string
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

		println(log)
		


	})
}