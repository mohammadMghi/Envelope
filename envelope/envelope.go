package envelope

import (
	"log"
	"net"
	"sync"

	"net/http"
	"reflect"
)



type Envelope struct{
	c sync.Pool
 	eHanlders []http.Handler
	Router Router
	log Log
	Port string
}

type EHandler interface{}

func New(port string ) *Envelope{
	router := NewRouter()
	log := NewLog()
	initLog(port)
	return &Envelope{
		Router: *router,
		log: log,
		Port:  port,
	}
}

func initLog(port string){
	print("Envelope now running on : "  + GetOutboundIP().String() + ":" + port + "\n")
}
// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP
}

func (e *Envelope)addHandler(handler http.Handler){
	// Checks if middlware handler is a function
	if reflect.TypeOf(handler).Kind() != reflect.Func{
		panic("type must be a callable function")
	}

	e.eHanlders = append(e.eHanlders , handler )
}



func (e *Envelope)addHandlers(handler ...http.Handler){

	for _ , handler := range(handler){
		e.addHandler(handler)
	}

 
}

type Middleware func(http.HandlerFunc) http.HandlerFunc
 
 

func (l *Envelope) ServeHTTP(w http.ResponseWriter , req *http.Request){

	path := req.URL.Path
	method := req.Method

	handler := l.Router.getHandler(method, path)
 
	print(" path : " + path)
	print(" method :" + method)



	handler.ServeHTTP(w, req)
 
}


 

