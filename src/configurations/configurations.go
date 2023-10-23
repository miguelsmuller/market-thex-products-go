package configurations

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
    APIAddress    string
}

func LoadEnvVars() (*Config, error) {
	err := godotenv.Load("../settings/.env")
    if err != nil {
        return nil, err
    }

    cfg := &Config{
        APIAddress:  os.Getenv("API_ADDRESS"),
    }

    return cfg, nil
}
