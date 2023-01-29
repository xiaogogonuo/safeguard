package static

import "html/template"

var Template *template.Template

func init() {
	t, err := template.ParseFiles("static/html/template.html")
	if err != nil {
		panic(err)
	}
	Template = t
}
