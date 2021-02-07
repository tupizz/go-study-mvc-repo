package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
	"sparkbox.com.br/controllers/formdata"
	"sparkbox.com.br/utils"
	"sparkbox.com.br/views"
)

// CreateUserController - factory for our class
func CreateUserController() *UserController {
	return &UserController{
		SignUp: views.NewView("bootstrap", "views/sign-up.gohtml"),
	}
}

// UserController - Like a class of O.O. Languages
type UserController struct {
	SignUp *views.View
}

// SignUpPage - GET /sign-up
func (u *UserController) SignUpPage(w http.ResponseWriter, r *http.Request) {
	utils.Must(u.SignUp.Render(w, nil))
}

// CreateUser - POST /sign-up
func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	utils.Must(r.ParseForm())

	// ways for getting post values
	fmt.Println(r.PostForm["email"])
	fmt.Println(r.PostFormValue("password"))

	dec := schema.NewDecoder()
	form := formdata.SignupForm{}

	// Cause we're passing the pointer to form we'll update this variable
	utils.Must(dec.Decode(&form, r.PostForm))

	fmt.Fprintln(w, form)
}
