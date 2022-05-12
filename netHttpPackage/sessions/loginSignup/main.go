package main

import (
	// "fmt"
	"net/http"
	"text/template"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct{
	Username string
	Fullname string
	Age string
	Password []byte
}

var(
	sessionDb = map[string]string{} //session-username
	userDb = map[string]User{} //username-user
	tmp *template.Template
)

func init()  {
	tmp = template.Must(template.ParseGlob("*.html"))
}

func main()  {
	http.HandleFunc("/",index)
	http.HandleFunc("/login",login)
	http.HandleFunc("/signup",signup)
	http.HandleFunc("/logout",logout)

	http.ListenAndServe(":5500",nil)
}

func index(w http.ResponseWriter,r *http.Request)  {
	u := getUser(w,r)
	if u.Username=="" {
		tmp.ExecuteTemplate(w,"index.html",nil)
		return
	}
	tmp.ExecuteTemplate(w,"index.html",u)
}

func login(w http.ResponseWriter,r *http.Request)  {
	if alreadyLoggedIn(r){
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost{
		un := r.FormValue("username")
		pass := r.FormValue("password")

		//check username in db
		_,ok := userDb[un] 
		if !ok{
			http.Error(w,"Username or password do not match",http.StatusForbidden)
			http.Redirect(w,r,"/",http.StatusSeeOther)
			return
		}

		//check password in db
		passVal := userDb[un]
		err:=bcrypt.CompareHashAndPassword(passVal.Password,[]byte(pass))
		if err!= nil{
			http.Error(w,"Username or password do not match",http.StatusForbidden)
			http.Redirect(w,r,"/",http.StatusSeeOther)
			return
		}
		
		sId := uuid.NewV4()
		c := &http.Cookie{
			Name: "session",
			Value: sId.String(),
		}
		http.SetCookie(w,c)

		sessionDb[c.Value] = un
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
	
	tmp.ExecuteTemplate(w,"login.html",nil)
}

func signup(w http.ResponseWriter,r *http.Request)  {
	if alreadyLoggedIn(r){
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}

	var u User
	if r.Method == http.MethodPost{
		un := r.FormValue("username")
		fn := r.FormValue("fullname")
		age := r.FormValue("age")
		pass := r.FormValue("password")

		//checking if user already exists
		if _,ok := userDb[un]; ok{
			http.Error(w,"Username already taken",http.StatusForbidden)
			http.Redirect(w,r,"/signup",http.StatusSeeOther)
			return
		}

		//creating session
		sId := uuid.NewV4()
		c := &http.Cookie{
			Name: "session",
			Value: sId.String(),
		}
		http.SetCookie(w,c)

		//creating password hash
		passEnc,err := bcrypt.GenerateFromPassword([]byte(pass),bcrypt.MinCost)
		if err!=nil{
			http.Error(w,"Internal server error",http.StatusInternalServerError)
			return
		}

		//storing data in database
		u = User{un,fn,age,passEnc}
		sessionDb[c.Value] = un
		userDb[un] = u

		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}

	tmp.ExecuteTemplate(w,"signup.html",nil)
}

func logout(w http.ResponseWriter,r *http.Request)  {
	if !alreadyLoggedIn(r){
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}

	c,_ := r.Cookie("session")
	delete(sessionDb,c.Value)
	c = &http.Cookie{
		Name: "session",
		Value: "",
		MaxAge: 0,
	}
	http.SetCookie(w,c)
	http.Redirect(w,r,"/",http.StatusSeeOther)

}