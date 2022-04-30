package main

import (
	"fmt"
	"net/http"
)

func wel(rp http.ResponseWriter,rq *http.Request)  {
	rp.Header().Set("Content-Type","text/html; charset=utf-8")
	fmt.Fprintln(rp,`<h1>Car:</h1>`)
	fmt.Fprintf(rp,`<img src="/car">`)
}

func car(rp http.ResponseWriter,rq *http.Request)  {
	http.ServeFile(rp,rq,"car.jpg")
}

func main()  {
	http.HandleFunc("/",wel)
	http.HandleFunc("/car",car)

	http.ListenAndServe(":5500",nil)
}