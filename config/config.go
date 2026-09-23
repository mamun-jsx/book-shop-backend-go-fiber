package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APPport     string
	DBhost      string
	JWTsecret   string
	DATABASEurl string
}

// ? function to take the value from env and assign into config
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// * Load the env file to project
func LoadEnv() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}
	return &Config{
		APPport:     getEnv("APP_PORT", "8080"),
		DBhost:      getEnv("DB_HOST", "localhost"),
		JWTsecret:   getEnv("JWT_SECRET", "very99834secret4134"),
		DATABASEurl: getEnv("DATABASE_URL", ""),
	}
}