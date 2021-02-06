package views

import "html/template"

func NewView(layout string, files ...string) *View {
	files = append(
		files,
		"views/layouts/footer.gohtml",
		"views/layouts/bootstrap.gohtml",
		"views/layouts/navbar.gohtml",
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
