package main

import (
	"fmt"
	"log"
	"net/http"
)

func set(w http.ResponseWriter,r *http.Request)  {
	http.SetCookie(w,&http.Cookie{
		Name:"my-cookie",
		Value: "this is a delicious cookie",
	})
	fmt.Fprintln(w,"cookie written")
	fmt.Fprintln(w,"you can also check for cookie in your browser inspect-application")


}

func read(w http.ResponseWriter,r *http.Request)  {
	c,err:= r.Cookie("my-cookie")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintln(w,"Your cookie - ",c)
}

func main()  {
	http.HandleFunc("/",set)
	http.HandleFunc("/read",read)

	http.ListenAndServe(":9090",nil)
}