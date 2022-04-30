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

type myType struct{
	S []int
	A []string
}

func main()  {
	s1 := 13
	s2 := 14

	s:=[]int{s1,s2}

	a1:="Gaurav"
	a2:="Sahil"
	a3:="Tanya"
	a4:="Roy"

	a:=[]string{a1,a2,a3,a4}

	myStruct := myType{
		S: s,
		A: a,
	}

	err := tmp.ExecuteTemplate(os.Stdout,"index.html",myStruct)
	if err!=nil{
		log.Fatalln(err)
	}

}