package app

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type TemplateConfiguration struct {
	Path       string
	LayoutPath string
}

type AppConfiguration struct {
	Template TemplateConfiguration
}

var Conf AppConfiguration

func LoadConfiguration() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Conf.Template = TemplateConfiguration{
		Path:       os.Getenv("TEMPLATE_PATH"),
		LayoutPath: os.Getenv("TEMPLATE_PATH") + "/layouts",
	}
}
