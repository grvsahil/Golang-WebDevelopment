package main

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter, r *http.Request) User {
	var u User
	c, err := r.Cookie("session")
	if err != nil {
		sId := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sId.String(),
		}
	}
	http.SetCookie(w, c)

	uId, err := db.Query("SELECT userId FROM sessions where sessionId=?", c.Value)
	if err != nil {
		return u
	}

	var userId string
	for uId.Next() {
		uId.Scan(&userId)
	}

	user, err := db.Query("SELECT * from entries where id=?", userId)
	if err != nil {
		return u
	}

	for user.Next() {
		user.Scan(&u.Id, &u.Name, &u.Age, &u.Organization,&u.Password)
	}

	return u
}

func alreadyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}

	uId, err := db.Query("SELECT userId FROM sessions where sessionId=?", c.Value)
	if err != nil {
		return false
	}

	var userId string
	for uId.Next() {
		uId.Scan(&userId)
	}

	if userId==""{
		return false
	}

	// _, err = db.Query("SELECT * from entries where id=?", userId)
	// if err != nil {
	// 	return false
	// }
	
	
	return true
}
