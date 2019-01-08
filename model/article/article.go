package article

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/russross/blackfriday"
	"github.com/thenguyenit/go-blog/app"
)

//Article type is structure for a aritcle
type Article struct {
	Title     string `json:"title"`
	URL       string
	Avatar    string `json:"avatar"`
	Excerpt   string `json:"excerpt"`
	Content   template.HTML
	CreatedAt string `json:"created_at"`
	Status    bool   `json:"status"`
}

//Articles type is list of Article
type Articles []Article

//All will scan all of article filee in Article Path
func All() ([]Article, error) {
	articlePath := app.Conf.Article.Path

	result := Articles{}

	folders, err := ioutil.ReadDir(articlePath)
	if err == nil {
		for _, fd := range folders {
			if fd.IsDir() {
				year := fd.Name()
				folderPath := articlePath + "/" + year
				files, err := ioutil.ReadDir(folderPath)
				if err == nil {
					for _, f := range files {
						reJSON := regexp.MustCompile("(.*).json")
						if reJSON.MatchString(f.Name()) {
							article, err := readJSON(folderPath + "/" + f.Name())
							if err == nil {
								slug := strings.Replace(f.Name(), ".json", "", -1)
								article.URL = app.Conf.BaseURL + "/" + year + "/" + slug
								result = append(result, article)
							}
						}
					}
				}
			}
		}
	}

	fmt.Println("All:", result)

	return result, nil
}

//Load will .json file and .md file of article and return Article type
func Load(year string, articleSlug string) (Article, error) {
	jsonPath := app.Conf.Article.Path + "/" + year + "/" + articleSlug + ".json"
	article, err := readJSON(jsonPath)
	if err == nil {
		mdPath := app.Conf.Article.Path + "/" + year + "/" + articleSlug + ".md"
		bs, err := ioutil.ReadFile(mdPath)
		if err == nil {
			output := blackfriday.MarkdownBasic(bs)
			article.Content = template.HTML(output)

			return article, nil
		}
	}

	fmt.Println("Load", article)

	return Article{}, err
}

//readJSON will read .json file and return a Article type
func readJSON(path string) (Article, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return Article{}, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var aritcle Article

	json.Unmarshal([]byte(byteValue), &aritcle)

	if aritcle.Avatar != "" {
		aritcle.Avatar = app.Conf.PublicURL + aritcle.Avatar
	}

	fmt.Println("readJSON:", aritcle)

	return aritcle, nil
}
