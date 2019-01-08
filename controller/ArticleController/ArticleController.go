package ArticleController

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thenguyenit/go-blog/app"
	"github.com/thenguyenit/go-blog/model/article"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data, err := article.Load(vars["year"], vars["slug"])
	if err == nil {
		app.Render(w, "default", "detail", data)
	} else {
		log.Fatalf(err.Error())
	}
}
