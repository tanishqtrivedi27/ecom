package config

import (
	"os"
	"github.com/joho/godotenv"

)

type Config struct {
	PublicHost string
	DBUser     string
	DBPasword  string
	DBName     string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		DBUser:     getEnv("DB_USER", "http://localhost"),
		DBPasword:  getEnv("DB_PASSWORD", "http://localhost"),
		DBName:     getEnv("DB_NAME", "http://localhost"),
	}
}
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
