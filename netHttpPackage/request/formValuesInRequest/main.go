package main

import (
	"log"
	"net/http"
	"text/template"
)

type myType int

// Request is just a struct which has many field like
// Request{
// 	Method string
// 	URL *url.URL
// 	.
// 	.
// 	Form url.Values
// 	PostForm url.Values
// 	.
// 	.
// 	.
// }

// Form contains the parsed form data, including both the URL
// field's query parameters and the POST form data.
// This field is only available after ParseForm is called hence we are calling ParseForm on *request

//PostForm contains the parsed data only from POST form data
//rest everything is same as Form


func (mT myType) ServeHTTP(rp http.ResponseWriter,rq *http.Request)  {
	//func (r *Request) ParseForm() error

	// ParseForm populates r.Form and r.PostForm.
	// For all requests, ParseForm parses the raw query from the URL and updates r.Form.

	err := rq.ParseForm()
	if err != nil {
		log.Fatalln(err)
	} 
	
	tmp.ExecuteTemplate(rp,"index.html",rq.Form)
	
	// rq.Form is the Form field in request which is of type url.Values
	// url.Values has signature like - type Values map[string][]string

	// so basically, rq.Form is of type map[string][]string

}

var tmp *template.Template

func init()  {
	tmp = template.Must(template.ParseGlob("*"))
}

func main()  {
	var myVar myType
	http.ListenAndServe(":8081",myVar)
}