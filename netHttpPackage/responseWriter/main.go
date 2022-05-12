package main

import (
	"fmt"
	"net/http"
)

type myType int

func (mT myType) ServeHTTP(rp http.ResponseWriter, rq *http.Request) {
	// A ResponseWriter interface is used by an HTTP handler to construct an HTTP response.
	// We are trying to set header in responseWriter interface
	// see documentation :- https://pkg.go.dev/net/http#ResponseWriter

	// type Header map[string][]string
	// func (h Header) Set(key, value string)

	//when content type is set to text/plain browser is using this header and interpreting
	// as plain text
	// rp.Header().Set("content-type:","text/plain; charset=UTF-8")

	//when content type is set to text/html see what happens
	rp.Header().Set("Content-Type", "text/html; charset=utf-8")
	rp.Header().Set("myKey", "This is my key")
	fmt.Fprintln(rp, "<h2>Hey, how you doin man?</h2>")
}

func main() {

	var myVar myType

	http.ListenAndServe(":5500", myVar)

}
