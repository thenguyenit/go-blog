package structures

import (
	"io/ioutil"
	"time"
)

const (
	ArticlePath = ""
)

//Article is loaded from json file
type Article struct {
	CreatedAt time.Time
	Excerpt   string
	Status    bool
}

func pagination() {
	files, err := ioutil.ReadDir()
}
