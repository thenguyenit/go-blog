package app

import (
	"html/template"
	"log"
	"net/http"
)

func Render(w http.ResponseWriter, layoutName string, viewName string, data interface{}) {

	layoutPath := Conf.Template.LayoutPath + "/" + layoutName + ".html"
	viewPath := Conf.Template.Path + "/" + viewName + ".html"

	t, err := template.ParseFiles(layoutPath, viewPath)
	if err != nil {
		log.Fatal(err)
	}

	t.ExecuteTemplate(w, layoutName, data)
}
