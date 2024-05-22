package tools

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Env struct {
	Key   string
	Value string
}

func InitEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Error("Error loading .env file")
	}
}

func GetHostUrl() string {
	url := os.Getenv("HOST_URL")

	return url
}

func GetUsername() string {
	url := os.Getenv("USERNAME")

	return url
}

func GetPassword() string {
	url := os.Getenv("PASSWORD")

	return url
}

func GetLicenceKey() string {
	url := os.Getenv("LICENCE_KEY")

	return url
}
