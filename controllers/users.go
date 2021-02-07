package controllers

import (
	"fmt"
	"net/http"

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
	fmt.Println("Recebi corretamente")
}
