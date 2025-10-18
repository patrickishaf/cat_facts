package lib

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	RequestTimeoutSeconds int
}

var config AppConfig

func LoadEnvVariables() {
	godotenv.Load()
	timeoutStr := os.Getenv("REQUEST_TIMEOUT_SECONDS")
	timeout, err := strconv.Atoi(timeoutStr)
	if err != nil {
		timeout = 5
	}
	config = AppConfig{
		RequestTimeoutSeconds: timeout,
	}
}
