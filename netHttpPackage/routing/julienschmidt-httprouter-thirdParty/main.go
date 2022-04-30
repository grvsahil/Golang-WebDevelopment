package main

import (
	"net/http"
	"text/template"
	"github.com/julienschmidt/httprouter"
)

var tmp *template.Template

func init()  {
	tmp = template.Must(template.ParseGlob("*"))
}

func index(rp http.ResponseWriter,rq *http.Request,_ httprouter.Params)  {
	tmp.ExecuteTemplate(rp,"index.html",nil)
}

func about(rp http.ResponseWriter,rq *http.Request,_ httprouter.Params)  {
	tmp.ExecuteTemplate(rp,"about.html",nil)
}

func contact(rp http.ResponseWriter,rq *http.Request,_ httprouter.Params)  {
	tmp.ExecuteTemplate(rp,"contact.html",nil)
}

func submit(rp http.ResponseWriter,rq *http.Request,_ httprouter.Params)  {
	tmp.ExecuteTemplate(rp,"submit.html",nil)
}

func submitted(rp http.ResponseWriter,rq *http.Request,_ httprouter.Params)  {
	tmp.ExecuteTemplate(rp,"submitted.html",nil)
}

func user(rp http.ResponseWriter,rq *http.Request,this httprouter.Params)  {
	tmp.ExecuteTemplate(rp,"user.html",this.ByName("name"))
}

func main()  {
	mux := httprouter.New()
	mux.GET("/",index)
	mux.GET("/index",index)
	mux.GET("/about",about)
	mux.GET("/contact",contact)
	mux.GET("/submit",submit)
	mux.POST("/submit",submitted)
	mux.GET("/user/:name",user)

	http.ListenAndServe(":5500",mux)
}