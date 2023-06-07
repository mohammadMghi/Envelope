package envelope

import (
	 
	"log"
	"net"
	"strings"
	"sync"

	"net/http"
 
	"reflect"
)

type Envelope struct {
	c             sync.Pool
	HandlersChain []HandlersChain
	Router        Router
 
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

func (l *Envelope) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var handler http.Handler
	path := req.URL.Path
	method := req.Method

	result, isHttpHandler := l.Router.getHandler(method, path).(*http.HandlerFunc)

	createdHandler := l.createHandlerFunc(l.Router.getHandler(method, path))

	if isHttpHandler {
		print("createdHandler")
		handler = result
	} else {
		print("createdHandler")
		handler = createdHandler
	}

	print(" path : " + path)
	print(" method :" + method)

	handler.ServeHTTP(w, req)

}
func signature(f interface{}) string {
    t := reflect.TypeOf(f)
    if t.Kind() != reflect.Func {
        return "<not a function>"
    }

    buf := strings.Builder{}
    buf.WriteString("func (")
    for i := 0; i < t.NumIn(); i++ {
        if i > 0 {
            buf.WriteString(", ")
        }
        buf.WriteString(t.In(i).String())
    }
    buf.WriteString(")")
    if numOut := t.NumOut(); numOut > 0 {
        if numOut > 1 {
            buf.WriteString(" (")
        } else {
            buf.WriteString(" ")
        }
        for i := 0; i < t.NumOut(); i++ {
            if i > 0 {
                buf.WriteString(", ")
            }
            buf.WriteString(t.Out(i).String())
        }
        if numOut > 1 {
            buf.WriteString(")")
        }
    }

    return buf.String()
}