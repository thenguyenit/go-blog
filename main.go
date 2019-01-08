package main

import (
	"log"
	"net/http"
	"os"

	"github.com/thenguyenit/go-blog/app"
	"github.com/thenguyenit/go-blog/controller"
)

func main() {
	//load application configuration
	app.LoadConfiguration()

	//process all the request
	controller.HandleRequest()

	//start listen and serve
	log.Fatal(http.ListenAndServe(":"+os.Getenv("HTTP_PORT"), nil))
}
