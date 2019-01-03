package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/thenguyenit/go-blog/app/http/request"
)

func main() {
	//load bootstrap application
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	request.Handler()

	//start listen and serve
	log.Fatal(http.ListenAndServe(":"+os.Getenv("HTTP_PORT"), nil))
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
