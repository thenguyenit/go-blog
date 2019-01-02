package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	//load bootstrap application
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//call http request handler
	http.HandleFunc("/", handler)

	//start listen and serve
	log.Fatal(http.ListenAndServe(":"+os.Getenv("HTTP_PORT"), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	router := initRouter()

	route, err := router.search(r.URL.Path)
	if err != nil {
		//do something here
		fmt.Fprint(w, route.view)

		//
		route.handler()
	}

	fmt.Fprint(w, r.URL.Path)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

// func viewHandler(w http.ResponseWriter, r *http.Request) {
// 	title := r.URL.Path[len("/view/"):]
// 	p, _ := loadPage(title)
// 	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
// }

// func loadPage(title string) (*Page, error) {
// 	filename := ArticlePath + "/" + title + ".txt"
// 	body, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Page{Title: title, Body: body}, nil
// }
