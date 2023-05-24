package envlope

import (
	"log"
	"net/http"
	"time"
)

type Logger struct{
	handler http.Handler
}

func NewEvnlopeLogger(handdler http.Handler) *Logger{
	return &Logger{handler: handdler}
}

func (l *Logger) ServeHTTP(w http.ResponseWriter , r *http.Request){
    start := time.Now()
    l.handler.ServeHTTP(w, r)
    log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}



