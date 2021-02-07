package views

import (
	"html/template"
	"net/http"
	"path/filepath"

	"sparkbox.com.br/utils"
)

// Configuring the listeting for the layouts
var (
	LayoutDir   string = "views/layouts/"
	TemplateExt string = ".gohtml"
)

func NewView(layout string, files ...string) *View {
	files = append(
		files,
		layoutFiles()..., // destruct from slice
	)

	// Like destructuring from javascript
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	// Create new type of View and return a pointer to the variable of type view
	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}

// Here we are implemeting a interface just by creating a method with the same name
func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	utils.Must(v.Render(w, nil))
}

// Render (v *View) works like the 'this' from O.O. languages
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

// return a slice of string representing the layout files used in our application
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}
