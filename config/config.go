package config

import (
	"os"

	"github.com/joho/godotenv"
)

// type aliasing
type (
	Config struct {
		App App
	}

	App struct {
		Name  string
		Url   string
		Stage string
	}
)

func LoadConfig(path string) Config {
	godotenv.Load(path)
	return Config{
		App: App{
			Name:  os.Getenv("APP_NAME"),
			Url:   os.Getenv("APP_URL"),
			Stage: os.Getenv("APP_STAGE"),
		},
	}
}
