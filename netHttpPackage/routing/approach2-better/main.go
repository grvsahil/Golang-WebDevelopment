//serve mux is used to handle requests coming from different paths
//if path is /dog and method is get and
//if path is /cat and method is post, both will send different response
//mux or multiplexer runs different codes for different request signatures

//here this is basic approach and we have to create different types to handle
//different requests
package main

import (
	"fmt"
	"net/http"
)

type hotDog int
type hotCat int

func (hD hotDog) ServeHTTP(rp http.ResponseWriter,rq *http.Request)  {
	fmt.Fprintln(rp,"Hey doggie")
}

func (hC hotCat) ServeHTTP(rp http.ResponseWriter,rq *http.Request)  {
	fmt.Fprintln(rp,"Hey cattie")
}

func main()  {
	var hD hotDog
	var hC hotCat


	mux := http.NewServeMux()

	//mux.Handle takes a path string and a handler for that path
	mux.Handle("/dog",hD)
	mux.Handle("/cat/",hC)
	//if we add forward slash after keyword then it handles any path written after
	//the keyword. e.g /cat/something/other is also handled by hC handler only
	//but /dog/something is not handled by hD because it is '/dog' only not '/dog/'

	http.ListenAndServe(":5500",mux)
}

// type ServeMux
// func NewServeMux() *ServeMux
// func (mux *ServeMux) Handle(pattern string, handler Handler)
// func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
// func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)
// func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)