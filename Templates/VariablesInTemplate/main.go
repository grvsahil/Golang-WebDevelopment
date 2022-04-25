//look at the index.html to see how a variable is created in template

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
	err := tmp.ExecuteTemplate(os.Stdout,"index.html",`Tanya`)
	if err!=nil{
		log.Fatalln(err)
	}
}