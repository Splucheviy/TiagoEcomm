package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Envs ...
var Envs = InitConfig()

// Config ...
type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPass                 string
	DBAddr                 string
	DBName                 string
	JWTExpirationInSeconds int64
	JWTSecret              string
}

// InitConfig ...
func InitConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost:             getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                   getEnv("PORT", "8080"),
		DBUser:                 getEnv("DB_USER", "root"),
		DBPass:                 getEnv("DB_PASSWORD", "password"),
		DBAddr:                 fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:                 getEnv("DB_NAME", "tiago_ecomm"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 3600*24*7),
		JWTSecret:              getEnv("JWT_Secret", "not-secret-anymore?"),
	}
}

// getEnv...
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
