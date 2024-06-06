package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	DBUser     string
	DBPasword  string
	DBName     string
	JWTExpirationInSeconds int64
	JWTSecret string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		DBUser:     getEnv("DB_USER", "http://localhost"),
		DBPasword:  getEnv("DB_PASSWORD", "http://localhost"),
		DBName:     getEnv("DB_NAME", "http://localhost"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 3600 * 24 * 7),
		JWTSecret: getEnv("JWT_SECRET", "friursgj509gevmtvt"),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, _ := strconv.ParseInt(value, 10, 64)
		return i
	}

	return fallback
}
