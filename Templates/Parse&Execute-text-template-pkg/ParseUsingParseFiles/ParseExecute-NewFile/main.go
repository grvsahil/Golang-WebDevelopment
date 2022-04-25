//Here we are trying to parse our index.html content to a new file using text/template package

package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main()  {
	//Parse content from index.html
	tmp,err := template.ParseFiles("index.html")
	if err!=nil{
		log.Fatalln(err)
	}

	//creates new file named newIndex.html
	nf , err := os.Create("newIndex.html")
	if err!=nil{
		fmt.Println("Unable to create file")
	}

	//writes template content to new file
	err = tmp.Execute(nf,nil)
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
