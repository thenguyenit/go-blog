package app

import (
	"fmt"
	"html/template"
	"net/http"
)

func Render(w http.ResponseWriter, layoutName string, templateName string, data interface{}) {

	layoutPath := Conf.Template.Path + "/" + templateName

	t, e := template.New(layoutName).ParseFiles(layoutPath)
	fmt.Println(e)
	t.Execute(w, data)

	fmt.Println(data)

	fmt.Fprintf(w, "NTN"+layoutPath)
}
