package main

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/satori/go.uuid"
)

type user struct {
	Username string
	Fullname string
	Age      string
}

var dbsession = map[string]string{} //session-username
var dbuser = map[string]user{}      //username-user

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseGlob("*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/greet",greet)
	http.Handle("/favicon.ico",http.NotFoundHandler())

	http.ListenAndServe(":9090", nil)
}

func greet(w http.ResponseWriter,r *http.Request)  {
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Fprintln(w,"Create user first, redirecting")
		http.Redirect(w,r,"/",http.StatusSeeOther)
	}

	un := dbsession[cookie.Value]
	u := dbuser[un]

	w.Header().Set("Content-Type","text/html; charset=utf-8")
	fmt.Fprintf(w,"<h2>Hey user how you doin?\nYour details are :- Name:- %s\nAge:- %s\nUserName:- %s\n</h2>",u.Fullname,u.Age,u.Username)
}

func index(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		id := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "session",
			Value: id.String(),
		}
		http.SetCookie(w, cookie)
	}

	var u user
	if un, ok := dbsession[cookie.Value]; ok {
		u = dbuser[un]
	}

	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		fl := r.FormValue("fullname")
		age := r.FormValue("age")
		u = user{un, fl, age}
		dbsession[cookie.Value] = un
		dbuser[un] = u
	}

	tmp.ExecuteTemplate(w,"index.html",u)

}
