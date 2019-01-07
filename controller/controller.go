package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thenguyenit/go-blog/app"
	"github.com/thenguyenit/go-blog/model/article"
)

func Handler() {

	router := mux.NewRouter()
	router.HandleFunc("/{year:[0-9]+}/{slug}", ArticleDetail)
	router.HandleFunc("/", IndexHandler)
	router.HandleFunc("/about-me", AboutHandler)

	http.Handle("/", router)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data := article.All()
	app.Render(w, "default", "index", data)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	data := article.Load("2016", "about-me")
	app.Render(w, "default", "detail", data)
}

func ArticleDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := article.Load(vars["year"], vars["slug"])
	app.Render(w, "default", "detail", data)
}
