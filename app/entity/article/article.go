package article

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

const (
	ArticlePath = "/Users/ntnguyen/go/src/github.com/thenguyenit/go-blog/storage/app/private/article"
)

//Article is loaded from json file
type Article struct {
	CreatedAt time.Time `json:"created_at"`
	Excerpt string `json:"excerpt"`
	Status  bool   `json:"status"`
	Content string
}

func Pagination() map[int32][]Article {
	var result = make(map[int32][]Article)
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
							iYear, _ := strconv.ParseInt(year, 10, 32)
							result[int32(iYear)] = append(result[int32(iYear)], article)
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
	fmt.Println(string(byteValue))
	return Article{
		CreatedAt:
	}
}
