package utils

import (
	"net/http"

	"github.com/gorilla/schema"
)

func ParseForm(r *http.Request, formAddrs interface{}) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	return schema.
		NewDecoder().
		Decode(formAddrs, r.PostForm)
}
