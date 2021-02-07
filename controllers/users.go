package controllers

import (
	"fmt"
	"net/http"

	"sparkbox.com.br/controllers/formdata"
	"sparkbox.com.br/utils"
	"sparkbox.com.br/views"
)

// CreateUserController - factory for our class
func CreateUserController() *UserController {
	return &UserController{
		SignUp: views.NewView("bootstrap", "views/users/sign-up.gohtml"),
	}
}

// UserController - Like a class of O.O. Languages
type UserController struct {
	SignUp *views.View
}

// SignUpView - GET /sign-up
func (u *UserController) SignUpView(w http.ResponseWriter, r *http.Request) {
	utils.Must(u.SignUp.Render(w, nil))
}

// CreateView - POST /sign-up
func (u *UserController) CreateView(w http.ResponseWriter, r *http.Request) {
	var form formdata.SignupForm

	// Passing the reference for the form because it will be update in the context
	utils.Must(utils.ParseForm(r, &form))
	fmt.Fprintln(w, form)
}
