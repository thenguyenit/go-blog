package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thenguyenit/go-blog/controller/AboutController"
	"github.com/thenguyenit/go-blog/controller/ArticleController"
	"github.com/thenguyenit/go-blog/controller/IndexController"
)

//HandleRequest will handler the web request base on a router
func HandleRequest() {

	//Define a web router
	router := mux.NewRouter()
	router.HandleFunc("/{year:[0-9]+}/{slug}", ArticleController.Handle)
	router.HandleFunc("/", IndexController.Handle)
	router.HandleFunc("/about-me", AboutController.Handle)

	//Handler the request base on the router
	http.Handle("/", router)

	//Handler all request type of static files (start with public) of the application
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./storage/app/public"))))
}
