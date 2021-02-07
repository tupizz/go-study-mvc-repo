package formdata

// struct tags - kind of metadata inside the struct
type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}
