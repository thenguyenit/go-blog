package article

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	blackfriday "gopkg.in/russross/blackfriday.v2"

	"github.com/thenguyenit/go-blog/app"
)

const (
	ArticlePath = "./storage/app/private/article"
)

//Article is loaded from json file
type Article struct {
	Title     string `json:"title"`
	Url       string
	Excerpt   string `json:"excerpt"`
	Content   string
	CreatedAt string `json:"created_at"`
	Status    bool   `json:"status"`
}

type Articles []Article

func All() []Article {
	result := Articles{}
	folders, err := ioutil.ReadDir(ArticlePath)
	if err == nil {
		for _, fd := range folders {
			if fd.IsDir() {
				year := fd.Name()
				folderPath := ArticlePath + "/" + year
				files, err := ioutil.ReadDir(folderPath)
				if err == nil {
					for _, f := range files {
						reJSON := regexp.MustCompile("(.*).json")
						if reJSON.MatchString(f.Name()) {
							article := readJSON(folderPath + "/" + f.Name())
							slug := strings.Replace(f.Name(), ".json", "", -1)
							article.Url = "/" + year + "/" + slug
							result = append(result, article)
						}
					}
				}
			}
		}
	}

	return result
}

func Load(year string, articleSlug string) Article {
	jsonPath := app.Conf.Article.Path + "/" + year + "/" + articleSlug + ".json"
	article := readJSON(jsonPath)

	mdPath := app.Conf.Article.Path + "/" + year + "/" + articleSlug + ".md"
	bs, err := ioutil.ReadFile(mdPath)
	if err != nil {

	}
	output := blackfriday.Run(bs)
	article.Content = string(output)

	return article
}

func readJSON(path string) Article {
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(jsonFile)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var aritcle Article

	json.Unmarshal([]byte(byteValue), &aritcle)

	return aritcle
}
