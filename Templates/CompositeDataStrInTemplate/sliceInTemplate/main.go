//find out index and index1

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
	fruits := []string{"Apple","Mango","Kiwi","Banana"}

	err:= tmp.ExecuteTemplate(os.Stdout,"index.html",fruits)
	if err!=nil{
		log.Fatalln(err)
	}

	err1:= tmp.ExecuteTemplate(os.Stdout,"index1.html",fruits)
	if err1!=nil{
		log.Fatalln(err)
	}



}