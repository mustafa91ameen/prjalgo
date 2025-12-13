package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost             string
	DBPort             string
	DBUser             string
	DBPassword         string
	DBName             string
	DBSSLMode          string
	ServerPort         string
	JWTSecret          string
	JWTExpiry          string
	RefreshTokenExpiry string
}

func Load() *Config {
	// Load .env file if it exists (try multiple paths)
	_ = godotenv.Load()
	_ = godotenv.Load("../../.env")

	return &Config{
		DBHost:             getEnv("DB_HOST", "localhost"),
		DBPort:             getEnv("DB_PORT", "5432"),
		DBUser:             getEnv("DB_USER", "postgres"),
		DBPassword:         getEnv("DB_PASSWORD", ""),
		DBName:             getEnv("DB_NAME", "app"),
		DBSSLMode:          getEnv("DB_SSLMODE", "disable"),
		ServerPort:         getEnv("SERVER_PORT", "8080"),
		JWTSecret:          getEnv("JWT_SECRET", ""),
		JWTExpiry:          getEnv("JWT_EXPIRY", "15m"),
		RefreshTokenExpiry: getEnv("REFRESH_TOKEN_EXPIRY", "168h"),
	}
}

func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSSLMode)
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
