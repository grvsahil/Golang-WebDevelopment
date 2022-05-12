package main

import (
	// "fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func getUser(w http.ResponseWriter, r *http.Request) User {
	var u User
	cookie, err := r.Cookie("token")
	if err != nil {
		return u
	}
	
	tokenStr := cookie.Value

	claims := &Claims{}

	_, _ = jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	userId := claims.Id

	user, err := db.Query("SELECT * from entries where id=?", userId)
	if err != nil {
		return u
	}

	for user.Next() {
		user.Scan(&u.Id, &u.Name, &u.Age, &u.Organization, &u.Password)
	}

	return u

}

func alreadyLoggedIn(r *http.Request) bool {
	cookie, err := r.Cookie("token")
	if err != nil {
		return false
	}

	tokenStr := cookie.Value

	if tokenStr == "" {
		return false
	}

	return true
}
