package config

import (
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	DB DBConfig
}

type DBConfig struct {
	Host string
	Port string
	User string
	Password string
	Database string
}

var cfg *config

func Load() {
	godotenv.Load(".env")
	cfg = new(config)
	cfg.DB = DBConfig{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	}
}

func GetDB() DBConfig	{
	return cfg.DB
}