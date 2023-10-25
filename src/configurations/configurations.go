package configurations

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
    APIAddress    string
}

func LoadEnvVars() (*Config, error) {
	godotenv.Load("../settings/.env")

    cfg := &Config{
        APIAddress:  os.Getenv("API_ADDRESS"),
    }

    return cfg, nil
}
