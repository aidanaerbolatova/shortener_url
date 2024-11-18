package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DBHost   string
	DBPort   string
	Password string
	User     string
	DBName   string
	SSLMode  string
	Host     string
	Port     string
}

func LoadConfig(path string) (Config, error) {

	if err := godotenv.Load(path); err != nil {
		return Config{}, err
	}

	cfg := Config{
		DBHost:   os.Getenv("DBHOST"),
		DBPort:   os.Getenv("DBPORT"),
		Password: os.Getenv("PASSWORD"),
		User:     os.Getenv("USERNAME"),
		DBName:   os.Getenv("DBNAME"),
		SSLMode:  os.Getenv("SSLMODE"),
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
	}

	return cfg, nil
}
