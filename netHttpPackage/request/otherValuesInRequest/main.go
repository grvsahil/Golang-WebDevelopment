package main

import (
	"log"
	"net/http"
	"net/url"
	"text/template"
)

var tmp *template.Template

type myType int



func init()  {
	tmp = template.Must(template.ParseGlob("*"))
}

func (mT myType) ServeHTTP(rp http.ResponseWriter,rq *http.Request)  {
	err := rq.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method string
		URL *url.URL
		Header http.Header
		Host string
		ContentLength int64
		Form url.Values
	}{
		Method: rq.Method,
		URL: rq.URL,
		Header: rq.Header,
		Host: rq.Host,
		ContentLength: rq.ContentLength,
		Form: rq.Form,
	}

	tmp.ExecuteTemplate(rp,"index.html",data)
}

func main()  {
	var myVar myType
	http.ListenAndServe(":5500",myVar)
}