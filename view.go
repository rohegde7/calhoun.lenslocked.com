package main

import "html/template"

func NewView(files ...string) *View {

	files = append(files, "views/layout/footer.gohtml")

	templ, err := template.ParseFiles(files...)

	if err != nil {
		panic(err)

	}

	return &View{
		Template: templ,
	}
}

type View struct {
	Template *template.Template
}
