// we can use {. | func1 | func2 | func3}} to pass value of . to func1 and then return of 
// func2 to func3.
// This way we establish a pipeline

package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tmp *template.Template

var fm = template.FuncMap{
	"Time":currentTime,
	"Date":currentDate,
	"DateinDay" : currentDateinDays,
}

func currentTime(t time.Time) string{
	return t.Format("03:04:05PM")
}

func currentDate(t time.Time) string{
	return t.Format("02/01/2006")
}

func currentDateinDays(t string) string{
	return t[0:2]
}

func init()  {
	tmp = template.Must(template.New("").Funcs(fm).ParseGlob("*"))
}

func main()  {
	err:= tmp.ExecuteTemplate(os.Stdout,"index.html",time.Now())
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