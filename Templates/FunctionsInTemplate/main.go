// We can use functions inside our index.html created in this main.go so that we can perform
// more operations on our Templates

// To pass a function a special type of data structure predefined for our use is FuncMap
// FuncMap is defined as type FuncMap map[string]any

package main

import (
	"log"
	"os"
	"text/template"
)

//we create variable fm of type func map
var fm = template.FuncMap{
	"trim":trimThree,
	"len" : length,
}

var tmp *template.Template

func init()  {
	//we try to assign template as we used to do

	// tmp = template.Must(template.ParseGlob("*"))

	//then we try to use funcs which looks like this on template pointer
	// func (t *Template) Funcs(funcMap FuncMap) *Template

	// tmp.Funcs(fm)

	//but this throws an error saying trim and len is not defined
	//this happens because we are trying to create a template pointer with the
	// trim and len written in index.html but those variables are not defined at this time
	// we are assigning it after this pointer is created so how will it know at the time
	// of creating template pointer from index.html

	//we have to define a pointer template with func at the very beginnning
	//we use    func New(name string) *Template    inside must to define an empty pointer template

	tmp = template.Must(template.New("").Funcs(fm).ParseGlob("*"))

	//template.must takes template pointer and error 
	//template.new("") gives an empty pointer
	//.func(fm) sets func map to this empty pointer and return new pointer
	//parsefiles return pointer and error which is used by template.must

}

func length(s string) int{
	return len(s)
}

//defining a function to use in our index.html
func trimThree(s string) string {
	return s[0:3]
}

func main()  {
	s1 := "Gaurav"
	s2 := "Tanya"
	s3 := "Shreya"
	s4 := "Mamta"

	s := []string{s1,s2,s3,s4}

	err := tmp.ExecuteTemplate(os.Stdout,"index.html",s)
	if err!=nil{
		log.Fatalln(err)
	}
}

//List of available functions in text/template package

// type Template

// func Must(t *Template, err error) *Template
// func New(name string) *Template
// func ParseFS(fsys fs.FS, patterns ...string) (*Template, error)
// func ParseFiles(filenames ...string) (*Template, error)
// func ParseGlob(pattern string) (*Template, error)

// func (t *Template) AddParseTree(name string, tree *parse.Tree) (*Template, error)
// func (t *Template) Clone() (*Template, error)
// func (t *Template) DefinedTemplates() string
// func (t *Template) Delims(left, right string) *Template
// func (t *Template) Execute(wr io.Writer, data any) error
// func (t *Template) ExecuteTemplate(wr io.Writer, name string, data any) error
// func (t *Template) Funcs(funcMap FuncMap) *Template
// func (t *Template) Lookup(name string) *Template
// func (t *Template) Name() string
// func (t *Template) New(name string) *Template
// func (t *Template) Option(opt ...string) *Template
// func (t *Template) Parse(text string) (*Template, error)
// func (t *Template) ParseFS(fsys fs.FS, patterns ...string) (*Template, error)
// func (t *Template) ParseFiles(filenames ...string) (*Template, error)
// func (t *Template) ParseGlob(pattern string) (*Template, error)
// func (t *Template) Templates() []*Template
