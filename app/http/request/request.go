package request

import (
	"fmt"
	"net/http"

	"github.com/thenguyenit/go-blog/app/entity/article"
)

func Handler() {
	//call http request handler
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/about/", AboutHandler)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Start pagination")
	list := article.Pagination()
	fmt.Println(list)
	for y, articles := range list {
		// 	fmt.Println(a)
		fmt.Fprintf(w, "Year"+string(y))
		for _, a := range articles {
			fmt.Fprintf(w, "Excert"+a.Excerpt)
		}

	}
	fmt.Println("End pagination")
	fmt.Fprint(w, r.URL.Path)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.Path)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
