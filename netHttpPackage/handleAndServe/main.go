package main

import (
	"fmt"
	"net/http"
)

type motorInfo string

func (mI motorInfo) ServeHTTP(rp http.ResponseWriter,rq *http.Request){
	fmt.Fprintln(rp,"This is served")
}
func main()  {
	var bmw motorInfo

	http.ListenAndServe(":8081",bmw)
}

// Any type which has ServeHTTP method with given parameters implements Handler interface
// So motorInfo becomes a Handler.

// ListenAndServe takes a TCP address i.e port num given here and a handler so we give our
// own handler

//Inside handler's ServeHTTP we can do anything
