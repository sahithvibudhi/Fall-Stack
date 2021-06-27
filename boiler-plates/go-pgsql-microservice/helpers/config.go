package helpers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	DBHost      string
	DBPort      string
	DBUser      string
	DBPass      string
	DBName      string
	HTTPPort    string
	Environment string
}

var Config Configuration

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		Log.Error(fmt.Sprintf("Error loading environmental varialbes from .env file. Err: %s", err.Error()))
		return err
	}

	Config.DBHost = os.Getenv("DB_HOST")
	Config.DBPort = os.Getenv("DB_PORT")
	Config.DBUser = os.Getenv("DB_USER")
	Config.DBPass = os.Getenv("DB_PASS")
	Config.DBName = os.Getenv("DB_NAME")
	Config.HTTPPort = os.Getenv("HTTP_PORT")

	return nil
}
