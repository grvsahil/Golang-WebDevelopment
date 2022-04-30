// Take the previous program in the previous folder and change it so that:
// a template is parsed and served
// you pass data into the template

package main

import (
	"net/http"
	"text/template"
)

func init()  {
	tmp = template.Must(template.ParseGlob("*"))	
}

var tmp *template.Template

func a(rp http.ResponseWriter,rq *http.Request)  {
	tmp.ExecuteTemplate(rp,"dog.html",rq.URL.Path)
}

func dog(rp http.ResponseWriter,rq *http.Request)  {
	tmp.ExecuteTemplate(rp,"dog.html",rq.URL.Path)	
}

func me(rp http.ResponseWriter,rq *http.Request)  {
	tmp.ExecuteTemplate(rp,"dog.html",rq.URL.Path)
}

func main()  {
	http.HandleFunc("/",a)
	http.HandleFunc("/dog/",dog)
	http.HandleFunc("/me/",me)

	http.ListenAndServe(":5500",nil)
}