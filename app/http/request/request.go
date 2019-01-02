package request

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.Path)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
