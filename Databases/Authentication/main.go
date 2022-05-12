package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var (
	tmp    *template.Template
	db     *sql.DB
	err    error
	jwtKey = []byte("secret_key")
)

type Credentials struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type Claims struct {
	Id string
	jwt.StandardClaims
}

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

	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/login", login).Methods("GET")
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/signup", signup).Methods("GET")
	r.HandleFunc("/signup", signupHandler).Methods("POST")
	r.HandleFunc("/logout", logoutHandler).Methods("GET")
	r.HandleFunc("/profile", middleware(profileHandler)).Methods("GET")
	r.HandleFunc("/greet", middleware(greet)).Methods("GET")
	r.HandleFunc("/age", middleware(age)).Methods("GET")
	http.ListenAndServe(":5500", r)
}

func middleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tokenStr := cookie.Value

		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(tokenStr, claims,
			func(t *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		f(w, r)
	}
}
