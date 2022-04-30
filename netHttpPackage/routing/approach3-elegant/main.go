package main

import (
	"fmt"
	"net/http"
)

func a(rp http.ResponseWriter,rq *http.Request){
	fmt.Fprintln(rp,"Hey doggie")
}

func b(rp http.ResponseWriter,rq *http.Request){
	fmt.Fprintln(rp,"Hey cattie")
}

func c(rp http.ResponseWriter,rq *http.Request){
	fmt.Fprintln(rp,"Kaisa hai bhai")
}

func main()  {
	//HandlerFunc converts ordinary function into http handler
	http.Handle("/dog/",http.HandlerFunc(a))
	
	http.HandleFunc("/cat/",b)
	http.HandleFunc("/",c)
	//we can pass any function which take ResponseWriter and *Request inside HandleFunc

	http.ListenAndServe(":5500",nil)
	//when handler is nil it runs default serve mux
}
//default serve mux uses these functions directly from http package

// func Handle(pattern string, handler Handler)
// func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
