package main

import (
	"fmt"
	"net/http"
	"os"
)

func wel(rp http.ResponseWriter,rq *http.Request)  {
	rp.Header().Set("Content-Type","text/html; charset=utf-8")

	fmt.Fprintln(rp,`<h1>Car</h1>`)
	fmt.Fprintln(rp,`<img src="/bike">`)

}

func bike(rp http.ResponseWriter,rq *http.Request)  {
	f,err := os.Open("bike.jpg")
	if err != nil {
		http.Error(rp,"file not found",404)
	}

	fi,err := f.Stat()
	if err != nil {
		http.Error(rp,"file not found",404)
	}

	http.ServeContent(rp,rq,fi.Name(),fi.ModTime(),f)
}

func main()  {
	http.HandleFunc("/",wel)
	http.HandleFunc("/bike",bike)

	http.ListenAndServe(":5500",nil)
}