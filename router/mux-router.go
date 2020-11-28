package router

import(
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

type muxRouter struct{}

var(
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(writer http.ResponseWriter, request *http.Request)){
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) POST(uri string, f func(writer http.ResponseWriter, request *http.Request)){
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (*muxRouter) SERVE(port string){
	log.Println("Listening to port ", port)
	http.ListenAndServe(port, muxDispatcher)
}