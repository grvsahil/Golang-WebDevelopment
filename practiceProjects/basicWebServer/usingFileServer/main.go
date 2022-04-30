package main

import (
	"fmt"
	"net/http"
)

type formData struct{
	name string
	age string
}

func formHandler(rp http.ResponseWriter,rq *http.Request)  {
	err:=rq.ParseForm()
	if err != nil {
		http.Error(rp,"404 not found",http.StatusNotFound)
	}

	fD := formData{
		name:rq.FormValue("name"),
		age:rq.FormValue("age"),
	}

	fmt.Fprintf(rp,"Name: %s\nAge: %s",fD.name,fD.age)
}

func main()  {
	fileServer:=http.FileServer(http.Dir("."))
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler)

	http.ListenAndServe(":5500",nil)
}