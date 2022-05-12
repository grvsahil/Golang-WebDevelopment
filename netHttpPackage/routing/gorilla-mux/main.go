package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func queryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w,"Your query was page %s of %s",page,title)
}

func welHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Welcome\nEnter query in form /books/{title}/page/{pageno.}")
}

func main()  {
	r := mux.NewRouter()
	r.HandleFunc("/",welHandler)
	r.HandleFunc("/books/{title}/page/{page}",queryHandler)

	err:=http.ListenAndServe(":9091",r)
	if err != nil {
		log.Fatalln(err)
	}
}