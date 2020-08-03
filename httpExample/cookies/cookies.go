package main

import (
	"net/http"
	"fmt"
	"encoding/base64"
	"time"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:	"firstCookie",
		Value:	"Go Web Programming",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:	"secondCookie",
		Value:	"Macadamia Nut",
		HttpOnly: true,
	}

	// Set cookie in the header
	w.Header().Set("Set-Cookie", c1.String())
	// Alternatively you can set it using SetCookie
	http.SetCookie(w, &c2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	// get cookies directly from the header
	h := r.Header["Cookie"]
	fmt.Fprintln(w, h)

	// or use the Cookie function to get named cookie
	c1, err := r.Cookie("firstCookie")
	if err != nil {
		fmt.Println(w, "firstCookie unavailable")
	}
	// you are able to get all cookies in a slice 
	cookies := r.Cookies()
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cookies)
}

func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Hello world!")
	c := http.Cookie{
		Name: "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
}

func showMessage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "No message found")
		}
	} else {
		// Reset the cookie
		rc := http.Cookie{
			Name: "flash",
			MaxAge: -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	// do flash message example for displaying a message if available
	http.HandleFunc("/set_message", setMessage)
	http.HandleFunc("/show_message", showMessage)
	server.ListenAndServe()
}
