package envelope

import (
	"fmt"
	"log"
	"net"
	"sync"

	"net/http"
	"reflect"
)



type Envelope struct{
	c sync.Pool
	HandlersChain []HandlersChain
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

func (e *Envelope)addHandler(handler HandlersChain){
	// Checks if middlware handler is a function
	if reflect.TypeOf(handler).Kind() != reflect.Func{
		panic("type must be a callable function")
	}

	e.HandlersChain = append(e.HandlersChain , handler )
}



func (e *Envelope)addHandlers(handler ...HandlersChain){

	for _ , handler := range(handler){
		e.addHandler(handler)
	}

 
}
type responseWriter struct{
	http.ResponseWriter
}

type Context struct{
	writer responseWriter
	*http.Request
}

func (w *responseWriter) Unwrap() http.ResponseWriter {
	return w.ResponseWriter
}

type HandlerFunc func(*Context)
type HandlersChain  HandlerFunc


type Middleware func(http.HandlerFunc) http.HandlerFunc
 
 

func (e *Envelope) Handler() http.Handler{
	return e
}

func (e *Envelope) createHandlerFunc(handler Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler
	})
}

type Handler interface{}

func validateHandler(handler Handler) {
	if reflect.TypeOf(handler).Kind() != reflect.Func {
		panic("evelope handler must be a callable func")
	}
}

func (l *Envelope) ServeHTTP(w http.ResponseWriter , req *http.Request){
	var handler http.Handler
	path := req.URL.Path
	method := req.Method




	result, isHttpHandler:=l.Router.getHandler(method, path).(*http.HandlerFunc)

	createdHandler := l.createHandlerFunc(l.Router.getHandler(method, path))

	if isHttpHandler {
		print("createdHandler")
		handler = result
	}else{
		print("createdHandler")
		handler = createdHandler
	}

 
 
	print(" path : " + path)
	print(" method :" + method)



	handler.ServeHTTP(w, req)
 
}


 

