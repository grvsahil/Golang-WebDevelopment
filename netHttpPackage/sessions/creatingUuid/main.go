package main

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		cookie,err := r.Cookie("session-id")
		if err != nil {
			id := uuid.NewV4()
			cookie := &http.Cookie{
				Name: "session-id",
				Value: id.String(),
				HttpOnly: true,
			}
			http.SetCookie(w,cookie)
			
		}
		fmt.Fprintln(w,cookie)
	})

	http.ListenAndServe(":9090",nil)
}
