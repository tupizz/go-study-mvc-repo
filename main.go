package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"sparkbox.com.br/views"
)

var (
	homeView    *views.View
	contactView *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Someone visit our handleHome")
	w.Header().Set("Content-Type", "text/html")

	err := homeView.Template.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Someone visit our handleContact")
	w.Header().Set("Content-Type", "text/html")

	err := contactView.Template.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1 style=\"color: red;\">Not found :(</h1>")
}

func main() {
	homeView = views.NewView("views/home.gohtml")
	contactView = views.NewView("views/contact.gohtml")

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)
}
