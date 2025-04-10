package generator

import (
	"os"
)

type GenConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func NewGenConfigFromEnv() GenConfig {
	return GenConfig{
		User:     os.Getenv("GEN_DB_USER"),
		Password: os.Getenv("GEN_DB_PASSWORD"),
		Host:     os.Getenv("GEN_DB_HOST"),
		Port:     os.Getenv("GEN_DB_PORT"),
		DBName:   os.Getenv("GEN_DB_NAME"),
	}
}
