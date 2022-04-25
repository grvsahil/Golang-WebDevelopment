//Created 4 files bogo.txt, dogo.txt, pogo.txt, gogo.txt

package main

import (
	"log"
	"os"
	"text/template"
)

func main()  {
	//Parse content from multiple files into one template container
	//To parse multiple files at once to the template variable we can also use ParseGlob
	tmp, err:=template.ParseGlob("*.txt")
	if err!=nil{
		log.Fatalln(err)
	}

	//Execute by default gives content of first file only 
	err = tmp.Execute(os.Stdout,nil)
	if err!=nil{
		log.Fatalln(err)
	}

	//ExecuteTemplate is used to specify file to get specific content
	err = tmp.ExecuteTemplate(os.Stdout,"pogo.txt",nil)
	if err!=nil{
		log.Fatalln(err)
	}

	err = tmp.ExecuteTemplate(os.Stdout,"dogo.txt",nil)
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
