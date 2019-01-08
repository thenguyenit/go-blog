package AboutController

import (
	"log"
	"net/http"

	"github.com/thenguyenit/go-blog/app"
	"github.com/thenguyenit/go-blog/model/article"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	data, err := article.Load("2016", "about-me")
	if err == nil {
		app.Render(w, "default", "detail", data)
	} else {
		log.Fatal(err)
	}
}
