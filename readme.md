# Simple blog with Golang

This is a simple blog that has written by Golang. It can show the article from the `.md` files under markdown format.

### Technical

##### MVC Pattern
**Controller**, **Model**: I separated the controller, model for every package corresponding with the Controller and Model name
**View**: Contains the HTML files with parsing template of `html/template` package

File structure
```
./
├── app
│   ├── bootstrap.go                //For initial the application as the template configuration, application's path
│   └── template.go                 //For rendering to HTML
│
├── controller                      //Controller
│   ├── AboutController         
│   │   └── AboutController.go      //Display "about me" article
│   ├── ArticleController
│   │   └── ArticleController.go    //Display detail of a article │   │   
│   ├── IndexController
│   │   └── IndexController.go      //Display all of the article
│   │
│   └── handler.go                  //Main Controller will handler all the request and find a particular Controller
│ 
├── model                   
│   └── article
│       └── article.go
│ 
└── view
│    ├── detail.html
│    ├── index.html
│    └── layout
│        └── default.html
│        
├── storage                         //For storage
│   └── app
│       ├── private
│       │   └── article
│       │       ├── 2016
│       │       │   ├── about-me.json   //For a article information
│       │       │   └── about-me.md     //Detail a article
│       │       └── 2018
│       │           ├── a-week-with-golang.json
│       │           └── a-week-with-golang.md
│       └── public
│           └── img
│               ├── article
│               │   └── 2018
│               │       ├── a-week-with-golang.jpg
│               │       └── concurrenct.gif
│               └── t-rex.png
│
├── main.go                         //Main application
├── readme.md                       //Wiki 
├── .env                            //The application configuration
```
##### Pakagese:

* [gorilla/mux](https://github.com/gorilla/mux): Using for router
* [russross/blackfriday](https://github.com/russross/blackfriday): Using to parse from markdown to HTML
* [joho/godotenv](github.com/joho/godotenv): Using for reading .env
  
### Clone and Install the dependency packages

```
go get github.com/thenguyenit/go-blog
```

```
go get github.com/gorilla/mux
go get github.com/russross/blackfriday
go get github.com/joho/godotenv
```
Will download, compile the resource into your $GOPATH directory.

### Try
- Change directory to the `go-blog` repository
```
cd ~/go/src/github.com/thenguyenit/go-blog
```

- Run the web with `go run`
```
go run main.go
```
- Open browser with the URL: [http://localhost:8080](http://localhost:8080)

### Codebase
A simple `main.go` file
```go
//main.go
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

	//start listen and serve (port default: 8080)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("HTTP_PORT"), nil))
}
```
