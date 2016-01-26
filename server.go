package gof

import (
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

type Server struct {
	Router  *mux.Router
	Negroni *negroni.Negroni
	Render  *render.Render
}

func NewServer(options ...render.Options) *Server {
	s := new(Server)
	s.Router = mux.NewRouter()
	s.Negroni = negroni.New()
	s.Render = render.New(options...)
	return s
}

func (this *Server) StrictSlash(val bool) *mux.Router {
	return this.Router.StrictSlash(val)
}

// middleware
func (this *Server) Use(handler negroni.Handler) {
	this.Negroni.Use(handler)
}

func (this *Server) UseFunc(handlerFunc func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)) {
	this.Negroni.UseFunc(handlerFunc)
}

func (this *Server) UseHandlerFunc(handlerFunc func(rw http.ResponseWriter, r *http.Request)) {
	this.Negroni.UseHandlerFunc(handlerFunc)
}

func (this *Server) UseHandler(handler http.Handler) {
	this.Negroni.UseHandler(handler)
}

// router
func (this *Server) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	return this.Router.HandleFunc(path, f)
}

func (this *Server) Handle(path string, handler http.Handler) *mux.Route {
	return this.Router.Handle(path, handler)
}

func (this *Server) Start(addr string) {
	l := log.New(os.Stdout, "[GOF] ", 0)
	l.Printf("listening on %s", addr)
	this.Negroni.UseHandler(this.Router)
	l.Fatal(http.ListenAndServe(addr, this.Negroni))
}
