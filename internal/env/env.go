package env

import (
	"os"

	"github.com/joho/godotenv"
)

const (
	OpenAiToken = "OPEN_AI_TOKEN"
	GinPort     = "GIN_PORT"
)

var loaded bool

func Get(key string) string {
	Load()
	return os.Getenv(key)
}

func Load() {
	if loaded {
		return
	}

	godotenv.Load()
	loaded = true
}
