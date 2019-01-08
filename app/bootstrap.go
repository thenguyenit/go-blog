package app

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//TemplateConfiguration type is Template Configuration
type TemplateConfiguration struct {
	Path       string
	LayoutPath string
}

//ArticleConfiguration type is Article Configuration
type ArticleConfiguration struct {
	Path string
}

//Configuration type is Application Configuration
type Configuration struct {
	BaseURL     string
	Template    TemplateConfiguration
	Article     ArticleConfiguration
	StoragePath string
	PrivatePath string
	PublicPath  string
	PublicURL   string
}

//Conf will contain configuration of the application
var Conf Configuration

//LoadConfiguration will read the .evn file and init some application's path
func LoadConfiguration() {
	//Read the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Conf.BaseURL = os.Getenv("BASE_URL")
	Conf.StoragePath = os.Getenv("STORAGE_PATH")
	Conf.PrivatePath = Conf.StoragePath + "/app/private"
	Conf.PublicPath = Conf.StoragePath + "/app/public"
	Conf.PublicURL = Conf.BaseURL + "/public"

	Conf.Template = TemplateConfiguration{
		Path:       os.Getenv("TEMPLATE_PATH"),
		LayoutPath: os.Getenv("TEMPLATE_PATH") + "/layout",
	}

	Conf.Article = ArticleConfiguration{
		Path: Conf.PrivatePath + "/article",
	}
}
