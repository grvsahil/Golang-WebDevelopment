package main

import (
	"log"
	"os"
	"text/template"
)

var tmp *template.Template

func init()  {
	tmp = template.Must(template.ParseGlob("*"))
}

func main()  {
	err := tmp.ExecuteTemplate(os.Stdout,"index.html",`Gaurav`)
	if err!=nil{
		log.Fatalln(err)
	}
}

//{{.}} is used to set data at any point in template