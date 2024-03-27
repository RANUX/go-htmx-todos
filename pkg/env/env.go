package env

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	// If we don't have an .env file exit
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		logrus.Error("No .env file found")
		return
	}
	godotenv.Load()
}
