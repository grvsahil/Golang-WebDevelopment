// ListenAndServe on port ":8080" using the default ServeMux.

// Use HandleFunc to add the following routes to the default ServeMux:

// "/" "/dog/" "/me/

// Add a func for each of the routes.

// Have the "/me/" route print out your name.

package main

import (
	"fmt"
	"net/http"
)

func a(rp http.ResponseWriter,rq *http.Request)  {
	fmt.Fprintln(rp,"<h1>Welcome</h1>")
	rp.Header().Set("Content-Type","text/html; charset=utf-8")
}

func dog(rp http.ResponseWriter,rq *http.Request)  {
	fmt.Fprintln(rp,"<h1>Hey doggie<h1>")
	rp.Header().Set("Content-Type","text/html; charset=utf-8")
}

func me(rp http.ResponseWriter,rq *http.Request)  {
	fmt.Fprintln(rp,"<h1>Gaurav</h1>")
	rp.Header().Set("Content-Type","text/html; charset=utf-8")
}

func main()  {
	http.HandleFunc("/",a)
	http.HandleFunc("/dog/",dog)
	http.HandleFunc("/me/",me)

	http.ListenAndServe(":5500",nil)
}