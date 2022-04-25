package main

import (
	"log"
	"os"
	"text/template"
)

var tmp *template.Template

type Car struct{
	Brand string
	Power int
	ZerotoHun int
}

func init()  {
	tmp = template.Must(template.ParseGlob("*"))
}

func main()  {
	myCar := Car{
		Brand:"Skoda",
		Power:180,
		ZerotoHun: 8}

	err := tmp.ExecuteTemplate(os.Stdout,"index.html",myCar)
	if err!=nil{
		log.Fatalln(err)
	}
}