//Created 4 files bogo.txt, dogo.txt, pogo.txt, gogo.txt

package main

import (
	"log"
	"os"
	"text/template"
)

//tmp is of type poiner to template in template package and has package level scope
var tmp *template.Template

//Runs once when program loads
func init()  {
	// func Must(t *Template, err error) *Template
	// func ParseGlob(pattern string) (*Template, error)

	//must takes two arguments and these arguments are exactly the arguments which is returned by glob
	//so we can use must because it does error checking for us and returns pointer to template

	tmp = template.Must(template.ParseGlob("*.txt"))
	//parse all the files with .txt , does error checking and returns pointer to template

}

func main()  {
	err := tmp.ExecuteTemplate(os.Stdout,"gogo.txt",nil)
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
