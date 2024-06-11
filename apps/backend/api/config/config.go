package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var appConfig *Config = &Config{}

type Config struct {
	Jwt JWTConfig
}

func LoadConfig() (*Config, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	expiresIn, err := strconv.Atoi(os.Getenv("JWT_EXPIRES_IN"))
	if err != nil {
		log.Fatalf("Invalid value for JWT_EXPIRES_IN: %v", err)
	}

	cfg := Config{
		Jwt: JWTConfig{
			JwtSecret: os.Getenv("JWT"),
			Type:      os.Getenv("JWT_TYPE"),
			ExpiresIn: expiresIn,
		},
	}

	appConfig = &cfg

	return &cfg, nil
}

func GetConfig() Config {
	return *appConfig
}
