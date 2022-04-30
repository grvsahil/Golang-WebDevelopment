package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmp *template.Template
type formData struct{
	Name string
	Age string
}

func init()  {
	tmp = template.Must(template.ParseGlob("*.html"))
}

func entryHandler(rp http.ResponseWriter,rq *http.Request){
	tmp.ExecuteTemplate(rp,"index.html",nil)
}

func formHandler(rp http.ResponseWriter,rq *http.Request){
	err:=rq.ParseForm()
	if err != nil {
		http.Error(rp,"404 file not found",http.StatusNotFound)
	}

	fD := formData{
		Name:rq.Form.Get("name"),
		Age:rq.Form.Get("age"),
	}

	tmp.ExecuteTemplate(rp,"form.html",fD)
	
}

func main()  {
	http.HandleFunc("/",entryHandler)
	http.HandleFunc("/form",formHandler)

	fmt.Println("Starting server at port 5500")
	http.ListenAndServe(":5500",nil)
}