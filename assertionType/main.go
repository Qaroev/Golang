package main

import (
	"net/http"
	log2 "log"
	"os"
	"fmt"
)

func main() {
	http.ListenAndServe("8080", New())
}

func New() http.Handler {
	mux := http.NewServeMux()
	log := log2.New(os.Stdout, "web", log2.LstdFlags)
	app := &app{mux, log}
	mux.HandleFunc("/foo", app.foo)
	return app
}

type app struct {
	mux *http.ServeMux
	log *log2.Logger
}

func (a *app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}

func (a *app) foo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request to foo")
}
