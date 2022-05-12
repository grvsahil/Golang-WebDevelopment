package main

import (
	"fmt"
	"net/http"
)

func main()  {
	http.HandleFunc("/",index)

	http.ListenAndServe(":8081",nil)
}

func index(w http.ResponseWriter,r *http.Request)  {
	r.Header.Set("Content-Type","text/html; charset=utf-8")
	fmt.Fprintln(w,"<h1>Hello World from docker container</h1>")
}