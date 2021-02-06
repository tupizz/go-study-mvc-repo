package views

import "html/template"

type View struct {
	Template *template.Template
}

func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml")

	// Like destructuring from javascript
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	// Create new type of View and return a pointer to the variable of type view
	return &View{
		Template: t,
	}
}
