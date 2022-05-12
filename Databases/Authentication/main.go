package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	tmp *template.Template
	db  *sql.DB
	err error
)

type User struct {
	Id           string
	Name         string
	Age          string
	Organization string
	Password     []byte
}

func init() {
	tmp = template.Must(template.ParseGlob("*.html"))

}

func main() {
	db, err = sql.Open("mysql", "root:Mobile@123@tcp(localhost:3306)/mysql?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Database connected")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.ListenAndServe(":5500", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	if u.Id == "" {
		tmp.ExecuteTemplate(w, "index.html", nil)
	} else {
		tmp.ExecuteTemplate(w, "index.html", u)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		password := r.FormValue("password")

		_, err := db.Query("SELECT id FROM entries where id=?", id)
		if err != nil {
			http.Error(w, "UserId or Password do not match", http.StatusForbidden)
			return
		}

		pass, err := db.Query("SELECT password FROM entries where id=?", id)
		if err != nil {
			http.Error(w, "UserId or Password do not match", http.StatusForbidden)
			return
		}
		var passS string
		for pass.Next() {
			pass.Scan(&passS)
		}
		bcrypt.CompareHashAndPassword([]byte(passS), []byte(password))

		//create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)

		query := fmt.Sprintf(`INSERT INTO sessions VALUES ("%s", "%s")`, c.Value, id)
		_, err = db.Exec(query)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusForbidden)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmp.ExecuteTemplate(w, "login.html", nil)

}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r){
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		name := r.FormValue("name")
		age := r.FormValue("age")
		organization := r.FormValue("organization")
		password := r.FormValue("password")

		_, err := db.Query("SELECT id FROM entries where id=?", id)
		if err != nil {
			http.Error(w, "UserId already taken", http.StatusForbidden)
			return
		}

		encPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusForbidden)
			return
		}

		query := fmt.Sprintf(`INSERT INTO entries VALUES ("%s", "%s", "%s", "%s","%s")`, id, name, age, organization, string(encPass))
		_, err = db.Exec(query)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusForbidden)
		}

		fmt.Fprintln(w, "Account successfully created")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmp.ExecuteTemplate(w, "signup.html", nil)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r){
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
	c,_ := r.Cookie("session")

	_, err := db.Exec("DELETE FROM sessions WHERE sessionId=?", c.Value)
	if err != nil {
		http.Error(w,"Internal server error",http.StatusForbidden)
		return
	}

	c = &http.Cookie{
		Name: "session",
		Value: "",
		MaxAge: -1,
	}
	http.SetCookie(w,c)

	http.Redirect(w,r,"/",http.StatusSeeOther)
}
