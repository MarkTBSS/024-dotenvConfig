package main

import (
	"log"
	"os"

	"github.com/MarkTBSS/24-dotenvConfig/config"
)

func main() {
	// Initialize config
	cfg := config.LoadConfig(func() string {
		// [0]main.go, [1]./env/xxx/file.ext
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		}
		return os.Args[1]
	}())
	// Check Config Object
	log.Println(cfg)
}
