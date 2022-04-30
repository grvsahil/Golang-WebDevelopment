package main

import (
	"fmt"
	"net/http"
)

func images(rp http.ResponseWriter,rq *http.Request){
	rp.Header().Set("Content-Type","text/html; charset=utf-8")

	fmt.Fprintln(rp,`<img src="bike.jpg">`)
	fmt.Fprintln(rp,`<img src="car.jpg">`)
}

func main()  {
	http.Handle("/",http.FileServer(http.Dir(".")))
	http.HandleFunc("/images",images)


	http.ListenAndServe(":5500",nil)
}