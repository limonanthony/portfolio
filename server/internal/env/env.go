package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/limonanthony/portfolio/internal/logger"
)

type Key string

const (
	HttpHost   Key = "HTTP_HOST"
	HttpPort   Key = "HTTP_PORT"
	HttpSecure Key = "HTTP_SECURE"
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		logger.Warnf("Error loading .env file %v", err)
	}
}

func Get(key Key) string {
	value := os.Getenv(string(key))
	if value == "" {
		logger.Panicf("Environment variable %s not set", string(key))
	}

	return value
}

func GetInt(key Key) int {
	value := Get(key)

	intVal, err := strconv.Atoi(value)
	if err != nil {
		logger.Panicf("Error converting %s to int", value)
	}

	return intVal
}

func GetBool(key Key) bool {
	value := Get(key)

	boolVal, err := strconv.ParseBool(value)
	if err != nil {
		logger.Panicf("Error converting %s to bool", value)
	}

	return boolVal
}
