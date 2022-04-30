package main

import (
	"io"
	"net/http"
	"os"
	"text/template"
)

var tmp *template.Template

func init()  {
	tmp =template.Must(template.ParseGlob("*.html"))
}

func welcome(rp http.ResponseWriter,rq *http.Request)  {
	tmp.ExecuteTemplate(rp,"index.html",nil)
}

func car(rp http.ResponseWriter,rq *http.Request)  {
	myCar := `<img src="/car/set">`
	tmp.ExecuteTemplate(rp,"index.html",myCar)
}

func setCar(rp http.ResponseWriter,rq *http.Request)  {
	f,err:=os.Open("car.jpg")
	if err != nil {
		http.Error(rp,"file not found, hehe",404)
	}
	defer f.Close()

	io.Copy(rp,f)
}

func bike(rp http.ResponseWriter,rq *http.Request)  {
	myCar := `<img src="/bike/set">`
	tmp.ExecuteTemplate(rp,"index.html",myCar)
}

func setBike(rp http.ResponseWriter,rq *http.Request)  {
	f,err:=os.Open("bike.jpg")
	if err != nil {
		http.Error(rp,"file not found, hehe",404)
	}
	defer f.Close()

	io.Copy(rp,f)
}


func main()  {
	http.HandleFunc("/",welcome)
	http.HandleFunc("/car",car)
	http.HandleFunc("/car/set",setCar)
	http.HandleFunc("/bike",bike)
	http.HandleFunc("/bike/set",setBike)

	http.ListenAndServe(":5500",nil)
}