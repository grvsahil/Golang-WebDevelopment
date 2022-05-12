package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmp.ExecuteTemplate(w, "index.html", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials

	credentials.Id = r.FormValue("id")
	credentials.Password = r.FormValue("password")

	_, err = db.Query("SELECT id FROM entries where id=?", credentials.Id)
	if err != nil {
		http.Error(w,"Username or Password do not match",http.StatusUnauthorized)
		return
	}

	pass, err := db.Query("SELECT password FROM entries where id=?", credentials.Id)
	if err != nil {
		http.Error(w,"Username or Password do not match",http.StatusUnauthorized)
		return
	}
	var passS string
	for pass.Next() {
		pass.Scan(&passS)
	}
	err = bcrypt.CompareHashAndPassword([]byte(passS), []byte(credentials.Password))
	if err != nil {
		http.Error(w,"Username or Password do not match",http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 4)

	claims := &Claims{
		Id: credentials.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}

func login(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		fmt.Fprintln(w, "Already logged in, logout first to login different user")
		return
	}

	tmp.ExecuteTemplate(w, "login.html", nil)

}

func signupHandler(w http.ResponseWriter, r *http.Request) {
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
}

func signup(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Error(w,"Already logged in, logout first to signup",http.StatusBadRequest)
		return
	}

	tmp.ExecuteTemplate(w, "signup.html", nil)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	cookie, _ := r.Cookie("token")

	cookie = &http.Cookie{
		Name:   "token",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	var u User
	u = getUser(w, r)
	tmp.ExecuteTemplate(w, "profile.html", u)
}

func greet(w http.ResponseWriter, r *http.Request)  {
	var u User
	u = getUser(w, r)
	tmp.ExecuteTemplate(w, "greet.html", u)
}

func age(w http.ResponseWriter, r *http.Request)  {
	var u User
	u = getUser(w, r)
	tmp.ExecuteTemplate(w, "age.html", u)
}
