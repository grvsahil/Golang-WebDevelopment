package main

import (
	"log"
	"os"
	"text/template"
)

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseGlob("*"))
}

func main() {
	myMap := map[int]string{
		1:  "Jan",
		2:  "Feb",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "Aug",
		9:  "Sep",
		10: "Oct",
		11: "Nov",
		12: "Dec"}

	err := tmp.ExecuteTemplate(os.Stdout,"index.html",myMap)
	if err!=nil{
		log.Fatalln(err)
	}
}
