package main

import (
	"fmt"
	"net/http"
)

type myType int

func (mT myType) ServeHTTP(rp http.ResponseWriter,rq *http.Request)  {
	switch rq.URL.Path {
	case "/dog/":
		fmt.Fprintln(rp,"Hey doggie")
	case "/cat/":
		fmt.Fprintln(rp,"Hey cattie")
	default:
		fmt.Fprintln(rp,"Kaisa hai bhai")	
	}
}

func main()  {
	var mT myType
	http.ListenAndServe(":5500",mT)
}