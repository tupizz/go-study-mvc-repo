package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Someone visit our handleHome")
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my site!</h1>")
}

func handleContact(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Someone visit our handleContact")
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get it touch, please send an email to <a href=\"mailto:tadeu.tupiz@gmail.com\">tadeu.tupiz@gmail.com</a>")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleHome)
	r.HandleFunc("/contact", handleContact)
	http.ListenAndServe(":3000", r)
}
