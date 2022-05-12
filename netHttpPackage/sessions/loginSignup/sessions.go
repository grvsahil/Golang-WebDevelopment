package main

import(
	"net/http"
	"github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter,r *http.Request) User {
	c,err := r.Cookie("session")
	if err != nil {
		sId := uuid.NewV4()
		c = &http.Cookie{
			Name: "session",
			Value: sId.String(),
		}
	}
	http.SetCookie(w,c)

	var u User
	if un,ok := sessionDb[c.Value]; ok{
		u = userDb[un]
	}
	return u

}

func alreadyLoggedIn(r *http.Request) bool {
	c,err := r.Cookie("session")
	if err != nil {
		return false
	}
	un := sessionDb[c.Value]
	_,ok := userDb[un]
	return ok
}