package controller

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/thenguyenit/go-blog/app"
	"github.com/thenguyenit/go-blog/model/article"
)

func Handler() {

	//call http request handler
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/about/", AboutHandler)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// data := article.All()
	layoutPath := app.Conf.Template.Path + "/" + "index.tmpl"

	t, e := template.New("index.tmpl").ParseFiles(layoutPath)
	fmt.Println(e)
	t.Execute(w, article.Article{
		Excerpt: "Hello 1",
	})

	// fmt.Println(data)

	app.Render(w, "master.tmpl", "index.tmpl", article.Article{
		Excerpt: "Hello",
	})
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.Path)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

// func loadPage(title string) (*Page, error) {
// 	filename := title + ".txt"
// 	body, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Page{Title: title, Body: body}, nil
// }
