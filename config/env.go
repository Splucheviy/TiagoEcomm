package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var Envs = InitConfig() // var for InitConfig

// Config...
type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPass     string
	DBAddr     string
	DBName     string
}

// InitConfig ...
func InitConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPass:     getEnv("DB_PASSWORD", "password"),
		DBAddr:     fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:     getEnv("DB_NAME", "tiago_ecomm"),
	}
}

// getEnv...
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
