package article

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"time"
)

const (
	ArticlePath = "./storage/app/private/article"
)

//Article is loaded from json file
type Article struct {
	Title     string `json:"title"`
	Excerpt   string `json:"excerpt"`
	Content   string
	CreatedAt time.Time
	Status    bool `json:"status"`
}

func All() map[string][]Article {
	var result = make(map[string][]Article)
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
							result[year] = append(result[year], article)
						}
					}
				}
			}
		}
	}

	return result
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

	fileInfo, err := jsonFile.Stat()
	if err == nil {
		aritcle.CreatedAt = fileInfo.ModTime()
	}

	return aritcle
}
