package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"sparkbox.com.br/controllers"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1 style=\"color: red;\">Not found :(</h1>")
}

func main() {
	staticController := controllers.CreateStaticController()
	cUsers := controllers.CreateUserController()

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)

	// Home
	r.Handle("/", staticController.HomeView).Methods("GET")

	// Contact
	r.Handle("/contact", staticController.ContactView).Methods("GET")

	// Users
	r.HandleFunc("/sign-up", cUsers.SignUpView).Methods("GET")
	r.HandleFunc("/sign-up", cUsers.CreateView).Methods("POST")

	http.ListenAndServe(":3000", r)
}
