// Take the previous program and change it so that:
// func main uses http.Handle instead of http.HandleFunc
// Contstraint: Do not change anything outside of func main

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
	http.Handle("/",http.HandlerFunc(a))
	http.Handle("/dog/",http.HandlerFunc(dog))
	http.Handle("/me/",http.HandlerFunc(me))

	http.ListenAndServe(":5500",nil)
}