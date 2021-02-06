package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Someone visit our handleHome")
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my site!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Someone visit our handleContact")
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get it touch, please send an email to <a href=\"mailto:tadeu.tupiz@gmail.com\">tadeu.tupiz@gmail.com</a>")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1 style=\"color: red;\">Not found :(</h1>")
}

func main() {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)
}
