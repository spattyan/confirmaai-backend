package helper

import (
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	ServerPort string `validate:"required;min=2"`
	Dsn        string `validate:"required;min=10"`
	AuthToken  string `validate:"required;min=10"`
}

func SetupEnv() (config Environment, err error) {
	if os.Getenv("APP_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			return Environment{}, err
		}
	}

	return Environment{
		ServerPort: os.Getenv("HTTP_PORT"),
		Dsn:        os.Getenv("DSN"),
		AuthToken:  os.Getenv("AUTH_TOKEN"),
	}, nil
}
