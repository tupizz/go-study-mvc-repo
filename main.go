package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"sparkbox.com.br/utils"
	"sparkbox.com.br/views"
)

var (
	homeView    *views.View
	contactView *views.View
	signUpView  *views.View
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1 style=\"color: red;\">Not found :(</h1>")
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	utils.Must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	utils.Must(contactView.Render(w, nil))
}

func signUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	utils.Must(signUpView.Render(w, nil))
}

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	signUpView = views.NewView("bootstrap", "views/sign-up.gohtml")

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/sign-up", signUp)
	http.ListenAndServe(":3000", r)
}
