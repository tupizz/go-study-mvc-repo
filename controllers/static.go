package controllers

/**
 * Parei na aula 7.9
 */
import (
	"sparkbox.com.br/views"
)

// CreateStaticController - factory
func CreateStaticController() *StaticController {
	return &StaticController{
		HomeView:    views.NewView("bootstrap", "views/static/home.gohtml"),
		ContactView: views.NewView("bootstrap", "views/static/contact.gohtml"),
	}
}

type StaticController struct {
	HomeView    *views.View
	ContactView *views.View
}
