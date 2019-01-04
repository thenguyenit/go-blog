package main

import (
	"log"
	"net/http"
	"os"

	"github.com/thenguyenit/go-blog/app"
	"github.com/thenguyenit/go-blog/controller"
)

func main() {
	//load bootstrap application
	app.LoadConfiguration()

	//process all the
	controller.Handler()

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
