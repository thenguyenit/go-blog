package IndexController

import (
	"log"
	"net/http"

	"github.com/thenguyenit/go-blog/app"
	"github.com/thenguyenit/go-blog/model/article"
)

//Handle will a main process of IndexController
func Handle(w http.ResponseWriter, r *http.Request) {
	data, err := article.All()
	if err == nil {
		app.Render(w, "default", "index", data)
	} else {
		log.Fatalf(err.Error())
	}
}
