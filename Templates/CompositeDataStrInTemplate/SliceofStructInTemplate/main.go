package main

import (
	"log"
	"os"
	"text/template"
)

var tmp *template.Template

type Car struct {
	Brand     string
	Power     int
	ZerotoHun float32
}

func init() {
	tmp = template.Must(template.ParseGlob("*"))
}

func main() {
	myCar1 := Car{
		Brand:     "Skoda",
		Power:     180,
		ZerotoHun: 8}

	myCar2 := Car{
		Brand:     "BMW",
		Power:     230,
		ZerotoHun: 7.4}

	myCar3 := Car{
		Brand:     "Audi RS e-tron",
		Power:     368,
		ZerotoHun: 4.2}
	myCar4 := Car{
		Brand:     "Mclaren 720s",
		Power:     410,
		ZerotoHun: 3.8}

	myCars := []Car{myCar1,myCar2,myCar3,myCar4}	

	err := tmp.ExecuteTemplate(os.Stdout, "index.html", myCars)
	if err != nil {
		log.Fatalln(err)
	}
}
