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

//List of available functions in text/template package

package main

import (
	"log"
	"os"
	"text/template"
)

//First lets create a file to parse, lets make index.html only

func main()  {
	//template.ParseFiles return a pointer to template and error which we receive
	tmp,err := template.ParseFiles("index.html")
	if err!=nil{
		log.Fatalln(err)
	}

	//now we are executing the template by writing it to standard output
	//Execute takes pointer receiver of template type 
	err = tmp.Execute(os.Stdout,nil)
	if err!=nil{
		log.Fatalln(err)
	}

}