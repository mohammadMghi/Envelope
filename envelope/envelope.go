package envelope

import (
 
	"log"
	"net"

	"sync"

	"net/http"

	"reflect"
)

type Envelope struct {
	c             sync.Pool
	HandlersChain []HandlersChain
	Router        Router
	RouterGroup   RouterGroup
	log           Log
	Port          string

 
}

type EHandler interface{}

func New(port string) *Envelope {
	router := NewRouter()
	log := NewLog()
	initLog(port)
	return &Envelope{
		Router: *router,
		log:    log,
		Port:   port,
	}
}

func initLog(port string) {
	print("Envelope now running on : " + GetOutboundIP().String() + ":" + port + "\n")
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

func (e *Envelope) addHandler(handler HandlersChain) {
	// Checks if middlware handler is a function
	if reflect.TypeOf(handler).Kind() != reflect.Func {
		panic("type must be a callable function")
	}

	e.HandlersChain = append(e.HandlersChain, handler)
}

func (e *Envelope) addHandlers(handler ...HandlersChain) {

	for _, handler := range handler {
		e.addHandler(handler)
	}

}

type responseWriter struct {
	http.ResponseWriter
}

type Context struct {
	writer responseWriter
	*http.Request
}

func (w *responseWriter) Unwrap() http.ResponseWriter {
	return w.ResponseWriter
}

type HandlerFunc func(*Context)
type HandlersChain HandlerFunc

type Middleware func(http.HandlerFunc) http.HandlerFunc

func (e *Envelope) Handler() http.Handler {
	return e
}

func (e *Envelope) createHandlerFunc(handler Handler) http.Handler {
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlerStringFunction , casted:= handler.(func() string)
		handlerFunction , castedhandlerFunction:= handler.(func())
		handlerIntStringFunction , castedhandlerStringIntFunction:= handler.(func() (int, string) )
		if casted{
			 txt := handlerStringFunction()
			 b := []byte(txt)
			 w.Write(b)
			 print(txt)
			 return
		}
		if castedhandlerFunction{
			handlerFunction()
			return
		}
		if castedhandlerStringIntFunction{
		    code , txt := handlerIntStringFunction()
			b := []byte(txt)
			w.Write(b)
			w.WriteHeader(code)
			return
		}



	})
}

type Handler interface{}

func validateHandler(handler Handler) {

	
	if reflect.TypeOf(handler).Kind() != reflect.Func {
		panic("evelope handler must be a callable func")
	}
}

func (e *Envelope) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var handler http.Handler
	path := req.URL.Path
	method := req.Method

	isGroupPath := e.RouterGroup.checkPathIsGroup(GetRootGroupPath(path))

	if isGroupPath{

		result, isHttpHandler := e.RouterGroup.getHandlerGroup( path , method).(http.HandlerFunc)

 

		if isHttpHandler {
	 
			handler = result
		} else {
			createdHandler := e.createHandlerFunc(e.Router.getHandler(method, path))
	 
			handler = createdHandler
		}
	

		handler.ServeHTTP(w, req)

		return
	}
 
	result, isHttpHandler := e.Router.getHandler(method, path).(http.HandlerFunc)



	if isHttpHandler {
 
		handler = result
	} else {
		createdHandler := e.createHandlerFunc(e.Router.getHandler(method, path))
 
		handler = createdHandler
	}

 

	handler.ServeHTTP(w, req)

	return

}
 