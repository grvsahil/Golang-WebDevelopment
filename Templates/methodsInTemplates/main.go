package main

import (
	"log"
	"os"
	"text/template"
)

var tmp *template.Template

func init(){
	tmp = template.Must(template.ParseGlob("*"))
}

type myCar struct{
	Name string
	Power int
	TurboEngine bool
}

func (mC myCar) Torque() int {
	return mC.Power*2
}

func (mC myCar) PerformanceCar() bool{
	if mC.TurboEngine {
		return true
	}
	return false
} 

func (mC myCar) Brand(s string) string{
	return s[0:3]
}

func main()  {
	myCar1 := myCar{
		Name: "BMW X3",
		Power: 210,
		TurboEngine: true,
	}


	err := tmp.ExecuteTemplate(os.Stdout,"index.html",myCar1)
	if err!=nil{
		log.Fatalln(err)
	}
}